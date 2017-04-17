package httpmock

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	ms := New(t)
	require.Nil(t, ms.server, "for now, server must be nil")

}
