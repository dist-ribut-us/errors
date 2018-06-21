package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors(t *testing.T) {
	var err error = String("this is an error")
	assert.Error(t, err)

	err = Wrap("testing", err)
	assert.Equal(t, "testing -> this is an error", err.Error())
}
