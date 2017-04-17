package httpmock

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	ms := New(t, true, true)
	require.NotNil(t, ms.server, "server must be instantiated")

}

func TestGet(t *testing.T) {
	ms := New(t, true, true)

	resp, err := http.Get(ms.server.URL + "/test")
	require.NoError(t, err, "server are not fine :(")
	require.NoError(t, ClientDumpResponse(resp, true), "response are not ok :(")

	ms.server.Close()
}

func TestAddGetHandler(t *testing.T) {
	ms := New(t, false, false)
	ms.AddGetHandler("/test", func(w http.ResponseWriter, r *http.Request) {
		ServerDumpRequest(r, true)
	})

	resp, err := http.Get(ms.server.URL + "/test")
	require.NoError(t, err, "server are not fine :(")
	require.NoError(t, ClientDumpResponse(resp, true), "response are not ok :(")

	ms.server.Close()
}
