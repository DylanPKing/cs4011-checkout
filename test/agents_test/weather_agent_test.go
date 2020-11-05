package test

import (
	"../../src/agents"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_That_NewWeather_Creates_Struct_That_Is_Not_Nil(t *testing.T) {
	// Arrange
	dummySeed := rand.NewSource(1)

	// Act
	weatherAgent := agents.NewWeather(&dummySeed)

	// Assert
	assert.NotNil(t, weatherAgent)
}
