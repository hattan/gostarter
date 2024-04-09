package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	expected := "Hello World"
	result := getMessage()

	assert.Equal(t, expected, result, "Expected %s, but got %s", expected, result)
}
