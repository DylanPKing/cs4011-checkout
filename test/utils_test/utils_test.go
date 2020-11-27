// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

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
