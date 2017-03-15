package testing

import (
	"fmt"
	"testing"
	"time"

	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/jtopjian/gophertools"
)

func TestResourceRefresh(t *testing.T) {
	rrf := func() (interface{}, string, error) {
		return "x", "ACTIVE", nil
	}

	rr := gophertools.ResourceRefreshConf{
		PollInterval: 1 * time.Second,
		Timeout:      2 * time.Second,
		Target:       []string{"ACTIVE"},
		Refresh:      rrf,
	}

	_, err := rr.ResourceRefresh()

	th.CheckNoErr(t, err)
}

func TestResourceRefreshTimeout(t *testing.T) {
	rrf := func() (interface{}, string, error) {
		return "x", "BUILD", nil
	}

	rr := gophertools.ResourceRefreshConf{
		PollInterval: 1 * time.Second,
		Timeout:      2 * time.Second,
		Target:       []string{"ACTIVE"},
		Pending:      []string{"BUILD"},
		Refresh:      rrf,
	}

	_, err := rr.ResourceRefresh()

	th.AssertEquals(t, "A timeout occurred", err.Error())
}

func TestResourceRefreshError(t *testing.T) {
	rrf := func() (interface{}, string, error) {
		return "x", "BUILD", fmt.Errorf("Error has occurred")
	}

	rr := gophertools.ResourceRefreshConf{
		PollInterval: 1 * time.Second,
		Timeout:      2 * time.Second,
		Target:       []string{"ACTIVE"},
		Pending:      []string{"BUILD"},
		Refresh:      rrf,
	}

	_, err := rr.ResourceRefresh()

	th.AssertEquals(t, "Error has occurred", err.Error())
}

func TestResourceRefreshExceed(t *testing.T) {
	rrf := func() (interface{}, string, error) {
		time.Sleep(4 * time.Second)
		return "x", "BUILD", fmt.Errorf("Just wasting time")
	}

	rr := gophertools.ResourceRefreshConf{
		PollInterval: 1 * time.Second,
		Timeout:      2 * time.Second,
		Target:       []string{"ACTIVE"},
		Pending:      []string{"BUILD"},
		Refresh:      rrf,
	}

	_, err := rr.ResourceRefresh()

	th.AssertEquals(t, "A timeout occurred", err.Error())
}

func TestResourceRefreshBadState(t *testing.T) {
	rrf := func() (interface{}, string, error) {
		return "x", "ERROR", nil
	}

	rr := gophertools.ResourceRefreshConf{
		PollInterval: 1 * time.Second,
		Timeout:      2 * time.Second,
		Target:       []string{"ACTIVE"},
		Pending:      []string{"BUILD"},
		Refresh:      rrf,
	}

	_, err := rr.ResourceRefresh()

	th.AssertEquals(t, "Unexpected state: ERROR", err.Error())
}
