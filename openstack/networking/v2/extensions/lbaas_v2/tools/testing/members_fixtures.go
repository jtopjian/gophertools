package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// HandleMemberDeleteCustomCode sets up the test server to issue a custom response.
func HandleMemberDeleteCustomCode(t *testing.T, code int) {
	th.Mux.HandleFunc("/lbaas/pools/1234asdf/members/1234asdf", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			switch code {
			case 404:
				w.WriteHeader(http.StatusNotFound)
			default:
				fmt.Fprintf(w, "{}")
			}

		case "DELETE":
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

			switch code {
			case 404:
				w.WriteHeader(http.StatusNotFound)
			case 409:
				w.WriteHeader(http.StatusConflict)
			case 500:
				w.WriteHeader(http.StatusInternalServerError)
			default:
				w.WriteHeader(http.StatusNoContent)
			}
		}

	})
}
