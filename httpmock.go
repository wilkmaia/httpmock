package httpmock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/fatih/color"
)

// MockServer wrapper for httptest.Server
type MockServer struct {
	server *httptest.Server
	t      *testing.T
}

func dumpServerRequest(r *http.Request, printBody bool) error {
	dump, err := httputil.DumpRequest(r, printBody)
	if err != nil {
		return err
	}
	begin := color.GreenString("server.request.begin")
	end := color.GreenString("server.request.end")
	fmt.Printf("[%s]\n%s\n[%s]\n", begin, dump, end)

	return nil
}

// New create a MockServer instance
func New(t *testing.T, dumpRequest bool, dumpBody bool) MockServer {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if true == dumpRequest {
			if err := dumpServerRequest(r, dumpBody); err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			}

		}
	}))

	return MockServer{
		server: server,
		t:      t,
	}
}
