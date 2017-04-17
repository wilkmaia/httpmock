package httpmock

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	ms := New(t, true, true)
	require.NotNil(t, ms.server, "server must be instantiated")

}
