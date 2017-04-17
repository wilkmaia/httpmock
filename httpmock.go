package httpmock

import (
	"net/http/httptest"
	"testing"
)

type MockServer struct {
	server *httptest.Server
	t      *testing.T
}

func New(t *testing.T) MockServer {

	return MockServer{
		server: nil,
		t:      t,
	}
}
