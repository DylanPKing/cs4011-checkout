package test

import (
	"testing"

	"../../src/agents"
	"github.com/stretchr/testify/assert"
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
		DataLogger: &agents.Logger{
			OutputFile: "out.txt",
		},
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

func Test_averageUtlisationLoop(t *testing.T) {

	totalTimePerCheckout := make([]float64, 4)
	avgTimePerCheckout := make([]float64, 4)
	utilisation := make([]float64, 4)

	processor := agents.DataProcessor{
		CheckoutUsage:      make(chan *agents.CheckoutUsageData),
		AvgCheckoutUseTime: 0.0,
		DataLogger: &agents.Logger{
			OutputFile: "out.txt",
		},
	}

	go func() {
		for i := 0; i < len(totalTimePerCheckout); i++ {
			processor.CheckoutUsage <- &agents.CheckoutUsageData{
				CheckoutNum:             i,
				TimeSpent:               10.0,
				TotalCustomersProcessed: 5,
			}
		}
	}()

	for i := 0; i < len(totalTimePerCheckout); i++ {
		data := <-processor.CheckoutUsage
		processor.AverageUtilisationLoop(
			&totalTimePerCheckout, &avgTimePerCheckout, &utilisation, data,
		)
	}

	for i := range totalTimePerCheckout {
		assert.Equal(t, totalTimePerCheckout[i], 10.0)
		assert.Equal(t, avgTimePerCheckout[i], 2.0)
		assert.Equal(t, utilisation[i], 0.25)
	}
}

func Test_ComputeAverageUtilisation(t *testing.T) {
	processor := agents.DataProcessor{
		CheckoutUsage:      make(chan *agents.CheckoutUsageData),
		AvgCheckoutUseTime: 0.0,
		DataLogger: &agents.Logger{
			OutputFile: "out.txt",
		},
	}

	go func() {
		for i := 0; i < 10; i++ {
			processor.CheckoutUsage <- &agents.CheckoutUsageData{
				CheckoutNum:             i,
				TimeSpent:               10.0,
				TotalCustomersProcessed: 5,
			}
		}
		close(processor.CheckoutUsage)
	}()

	go func() {
		processor.ComputeAverageUtilisation()
		_, ok := <-processor.CheckoutUsage
		assert.Equal(t, ok, false)
	}()
}

func Test_IncrementLostCustomers(t *testing.T) {
	processor := agents.DataProcessor{
		DataLogger: &agents.Logger{},
	}
	processor.IncrementLostCustomers()

	assert.Equal(t, processor.LostCustomers, int64(1))
}

func assertCheckoutUsageData(
	t *testing.T, actual *agents.CheckoutUsageData,
	num int, spent float64, processed int,
) {
	assert.Equal(t, actual.CheckoutNum, num)
	assert.Equal(t, actual.TimeSpent, spent)
	assert.Equal(t, actual.TotalCustomersProcessed, processed)
}
