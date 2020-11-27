// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

package test

import (
	"math/rand"
	"testing"

	"../../src/agents"
	"github.com/stretchr/testify/assert"
)

// Test_That_NewWeather_Creates_Struct_That_Is_Not_Nil ensures that the struct is created correctly and is not nil
func Test_That_NewWeather_Creates_Struct_That_Is_Not_Nil(t *testing.T) {
	// Arrange
	dummySeed := rand.NewSource(1)
	dummyDataProcessor := agents.DataProcessor{}

	// Act
	weatherAgent := agents.NewWeather(&dummySeed, &dummyDataProcessor)

	// Assert
	assert.NotNil(t, weatherAgent)
}

// Test_That_newCondition_Creates_Struct_That_Is_Not_Nil ensures that the struct is created correctly and is not nil
func Test_That_newCondition_Creates_Struct_That_Is_Not_Nil(t *testing.T) {
	// Arrange
	// Act
	testCondition := agents.NewCondition()

	// Assert
	assert.NotNil(t, testCondition)
}

// Test_setConditions_Sets_A_Set_Of_Conditions ensures that the setConditions sets the struct's condition set correctly
func Test_setConditions_Sets_A_Set_Of_Conditions(t *testing.T) {
	// Arrange
	// Two seeds with the same source return the same sequence of numbers
	dummySeed := rand.NewSource(1)
	dummySeed2 := rand.NewSource(1)
	dummyDataProcessor := agents.DataProcessor{}

	// Act
	weatherAgent1 := agents.NewWeather(&dummySeed, &dummyDataProcessor)
	weatherAgent2 := agents.NewWeather(&dummySeed2, &dummyDataProcessor)

	// Assert
	assert.Equal(t, weatherAgent1.Conditions, weatherAgent2.Conditions)
}

// Test_ToggleWeather_Toggles_Weather tests that the weather is successfully toggled using a deterministic number generator to mock the process
func Test_ToggleWeather_Toggles_Weather(t *testing.T) {
	// Arrange
	// Two seeds with the same source return the same sequence of numbers
	dummySeed := rand.NewSource(1)
	dummySeed2 := rand.NewSource(1)
	dummyLogger := agents.Logger{}
	dummyDataProcessor := agents.DataProcessor{DataLogger: &dummyLogger}

	// Act
	weatherAgent1 := agents.NewWeather(&dummySeed, &dummyDataProcessor)
	weatherAgent2 := agents.NewWeather(&dummySeed2, &dummyDataProcessor)
	weatherAgent1.ToggleWeather()

	// Assert
	assert.NotEqual(t, weatherAgent1.CurrentCondition, weatherAgent2.CurrentCondition)
}

// Test_ToggleWeather_Increments_TimesChangedToday tests that the TimesChangedToday is incremented when the weather is toggled
func Test_ToggleWeather_Increments_TimesChangedToday(t *testing.T) {
	// Arrange
	// Two seeds with the same source return the same sequence of numbers
	dummySeed := rand.NewSource(1)
	dummySeed2 := rand.NewSource(1)
	dummyLogger := agents.Logger{}
	dummyDataProcessor := agents.DataProcessor{DataLogger: &dummyLogger}

	// Act
	weatherAgent := agents.NewWeather(&dummySeed, &dummyDataProcessor)
	weatherAgent2 := agents.NewWeather(&dummySeed2, &dummyDataProcessor)
	weatherAgent.ToggleWeather()

	// Assert
	assert.NotEqual(t, weatherAgent.TimesChangedToday, weatherAgent2.TimesChangedToday)
	assert.Equal(t, 1, weatherAgent.TimesChangedToday)
}

// Test_That_ToggleWeather_Only_Toggles_TimesChangedLimit_Times test that toggling the weather can only occur as many times as TimesChangedLimit allows
func Test_That_ToggleWeather_Only_Toggles_TimesChangedLimit_Times(t *testing.T) {
	// Arrange
	dummySeed := rand.NewSource(1)
	dummyLogger := agents.Logger{}
	dummyDataProcessor := agents.DataProcessor{DataLogger: &dummyLogger}

	// Act
	weatherAgent := agents.NewWeather(&dummySeed, &dummyDataProcessor)
	// Deliberatelly iterating one more time than the limit to prove that the weather is not toggled an extra time
	for i := 0; i < weatherAgent.TimesChangedLimit+1; i++ {
		weatherAgent.ToggleWeather()
	}
	// Assert
	assert.Equal(t, weatherAgent.TimesChangedLimit, weatherAgent.TimesChangedToday)
}
