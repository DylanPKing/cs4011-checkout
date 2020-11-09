package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"../../src/agents"
)

func Test_CheckoutUsageData_Creation(t *testing.T) {
	actualValue := agents.CheckoutUsageData{
		CheckoutNum:             0,
		TimeSpent:               0.0,
		TotalCustomersProcessed: 0,
	}

	assertCheckoutUsageData(t, &actualValue, 0, 0.0, 0)
}

func Test_DataProcessor_Creation(t *testing.T) {
	actualValue := agents.DataProcessor{
		CheckoutUsage:      make(chan *agents.CheckoutUsageData),
		AvgCheckoutUseTime: 0.0,
	}

	go func() {
		actualValue.CheckoutUsage <- &agents.CheckoutUsageData{
			CheckoutNum:             0,
			TimeSpent:               0.0,
			TotalCustomersProcessed: 0,
		}
	}()

	actualCheckoutUsage := <-actualValue.CheckoutUsage

	assertCheckoutUsageData(t, actualCheckoutUsage, 0, 0.0, 0)

	assert.Equal(t, actualValue.AvgCheckoutUseTime, 0.0)
}

func assertCheckoutUsageData(
	t *testing.T, actual *agents.CheckoutUsageData,
	num int, spent float64, processed int,
) {
	assert.Equal(t, actual.CheckoutNum, num)
	assert.Equal(t, actual.TimeSpent, spent)
	assert.Equal(t, actual.TotalCustomersProcessed, processed)
}
