// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

package test

import (
	"testing"

	"../../src/manager"
	"github.com/stretchr/testify/assert"
)

// Test_New_Manager_Creates_Empty_Struct checks that calling the NewManager() functions returns an empty string.
func Test_New_Manager_Creates_Empty_Struct(t *testing.T) {
	//Arrange
	storeManager := manager.NewManager()

	//Assert
	assert.Equal(t, 0, storeManager.InitialNumberOfCheckouts)
	assert.Equal(t, 0, storeManager.NumberOfExpressCheckouts)
}
