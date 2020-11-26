package test

import (
	"math/rand"
	"testing"

	"../../src/agents"
	"github.com/stretchr/testify/assert"
)

func Test_That_RandnumGen_Generates_A_Random_Number_With_Parameter_Given(t *testing.T) {
	//Arrange
	dummySeed := rand.NewSource(1)
	//Act
	randomNum := agents.RandnumGen(&dummySeed, 10)
	//Assert
	assert.NotNil(t, randomNum)
}

func Test_That_NewCustomer_Creates_A_New_Customer(t *testing.T) {
	//Arrange
	dummySeed := rand.NewSource(1)
	dummyDataProcessor := agents.DataProcessor{}
	dummpyWeather := agents.Weather{}
	//Act
	aCustomer := agents.NewCustomer(&dummySeed, &dummyDataProcessor, &dummpyWeather)
	//Assert
	assert.NotNil(t, aCustomer)
}

func Test_That_NewProduct_Creates_A_New_Product(t *testing.T) {
	//Arrange
	dummySeed := rand.NewSource(1)
	//Act
	aProduct := agents.NewProduct(&dummySeed)
	//Assert
	assert.NotNil(t, aProduct)
}

func Test_That_FillTrolley_Fills_A_Trolley_With_Random_Number_Of_Items_Of_Random_Weight(t *testing.T) {
	//Arrange
	dummySeed := rand.NewSource(1)
	//Act
	aTrolley := agents.FillTrolley(&dummySeed)
	//Assert
	assert.NotNil(t, aTrolley)
}

func Test_That_ToggleQueue_Sets_Flag_For_Customer_Queueing(t *testing.T) {
	//Arrange
	dummySeed := rand.NewSource(1)
	dummyDataProcessor := agents.DataProcessor{}
	dummpyWeather := agents.Weather{}
	//Expected
	testQueue := true
	//Actual
	aCustomer := agents.NewCustomer(&dummySeed, &dummyDataProcessor, &dummpyWeather)
	aCustomer.ToggleQueue()
	assert.Equal(t, testQueue, aCustomer.Queue)
}
