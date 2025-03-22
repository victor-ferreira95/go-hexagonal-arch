package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_jsonError(t *testing.T) {
	msg := "hello json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"hello json"}`), result)
}
