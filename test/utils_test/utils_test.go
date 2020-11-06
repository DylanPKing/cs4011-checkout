package test

import (
	"testing"

	"../../src/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Sum_Success(t *testing.T) {
	slice := []float64{1, 2, 3, 4, 5}
	expectedValue := (float64)(1 + 2 + 3 + 4 + 5)
	actualValue := utils.Sum(&slice)

	assert.Equal(t, expectedValue, actualValue)
}
