package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlert(t *testing.T) {

	assert := require.New(t)

	a := NewAlert(AlertP0, "testSvc", "type1")

	assert.Equal(a.Priority(), AlertP0)
	assert.Equal(a.String(), "P0::testSvc::type1")
}

func TestAdd(t *testing.T) {

	assert := require.New(t)

	assert.Equal(add(1, 2), 3)
	assert.Equal(add(-1, 2), 1)
	assert.Equal(add(0, 2), 2)
}
