package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumber(t *testing.T) {
	got := RandomNumber(1, 5)
	assert.GreaterOrEqual(t, got, 1)
	assert.LessOrEqual(t, got, 5)
}
