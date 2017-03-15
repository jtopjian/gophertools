package gophertools

import (
	"fmt"
	"sync/atomic"
	"time"
)

// ResourceRefreshFunc is a function type that refreshes a resource to check if
// it has reached a certain state.
//
// It returns three values:
// resource: The resource being refreshed.
// state: the latest state of the resource.
// err: any errors that occurred.
//
// This is intentionally based on HashiCorp's Terraform implementation so
// ResourceRefreshFunc functions can be casted to Terraform's StateRefreshFunc.
type ResourceRefreshFunc func() (resource interface{}, state string, err error)

// ResourceRefreshConf is the configuration struct used for WaitForResource.
type ResourceRefreshConf struct {
	// Pending is valid states during a transition.
	Pending []string

	// PollInterval is the time between Refresh checks.
	PollInterval time.Duration

	// Refresh is a function to check the status of a resource.
	Refresh ResourceRefreshFunc

	// Target is the valid target state(s) of a resource.
	Target []string

	// Timeout is the amount of time to wait for a successful transition.
	Timeout time.Duration
}

// ResourceRefresh polls a ResourceRefreshFunc function.
// This is useful to wait for a resource to transition to a certain state.
// To handle situations when the ResourceRefreshFunc might hang indefinitely,
// the predicate will be prematurely cancelled after the timeout.
// It returns the latest resource and any errors that have occurred.
func (conf *ResourceRefreshConf) ResourceRefresh() (interface{}, error) {
	type ResourceRefreshResult struct {
		Resource interface{}
		State    string
		Error    error
	}
	var lastResult atomic.Value
	lastResult.Store(ResourceRefreshResult{})

	time.Sleep(1 * time.Second)

	ch := make(chan struct{})
	go func() {
		defer close(ch)

		for {
			resource, state, err := conf.Refresh()
			result := ResourceRefreshResult{
				Resource: resource,
				State:    state,
				Error:    err,
			}
			lastResult.Store(result)

			if err != nil {
				return
			}

			if resource == nil {
				result.Error = fmt.Errorf("Resource not found")
				lastResult.Store(result)
				return
			}

			for _, target := range conf.Target {
				if target == state {
					return
				} else {
					continue
				}
			}

			var found bool
			for _, pending := range conf.Pending {
				if pending == state {
					found = true
					break
				}
			}

			if !found {
				result.Error = fmt.Errorf("Unexpected state: %s", state)
				lastResult.Store(result)
				return
			}

			time.Sleep(conf.PollInterval)
		}
	}()

	select {
	case <-ch:
		r := lastResult.Load().(ResourceRefreshResult)
		return r.Resource, r.Error
	case <-time.After(conf.Timeout):
		r := lastResult.Load().(ResourceRefreshResult)
		return r.Resource, fmt.Errorf("A timeout occurred")
	}
}
