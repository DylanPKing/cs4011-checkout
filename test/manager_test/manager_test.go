package test

import (
	"../../src/manager"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_New_Manager_Creates_Empty_Struct checks that calling the NewManager() functions returns an empty string.
func Test_New_Manager_Creates_Empty_Struct(t *testing.T) {
	//Arrange
	storeManager := manager.NewManager()

	//Assert
	assert.Equal(t, 0, storeManager.InitialNumberOfCheckouts)
	assert.Equal(t, 0, storeManager.NumberOfExpressCheckouts)
}
