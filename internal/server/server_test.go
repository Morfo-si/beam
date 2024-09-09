package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	hello := "hello"
	assert.Equal(t, hello, "hello")
}
