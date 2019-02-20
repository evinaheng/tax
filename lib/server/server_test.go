package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tax/lib/server"
)

func TestGetIPAddress(t *testing.T) {
	assert.NotNil(t, GetIPAddress())
}
