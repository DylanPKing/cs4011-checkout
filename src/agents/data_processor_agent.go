package agents

import (
	"../utils"
)

// DataProcessor collects and performs calculations on data
// before sending data to the Logger
type DataProcessor struct {
	CheckoutUsage chan *CheckoutUsageData
	// CustomerData       chan *Customer
	AvgCheckoutUseTime float64
}

// ComputeAverageUtilisation collects the total usage of each checkout and
// computes their average.
func (processor *DataProcessor) ComputeAverageUtilisation() {
	totalTimePerCheckout := make([]float64, 10)
	avgTimePerCheckout := make([]float64, 10)
	utilisation := make([]float64, 10)
	for {
		data, ok := <-processor.CheckoutUsage
		if ok {
			processor.AverageUtilisationLoop(
				&totalTimePerCheckout, &avgTimePerCheckout, &utilisation, data,
			)
		}
	}
}

// AverageUtilisationLoop Is the core loop used to calculate the total, and
// average utilisation of each checkout.
func (processor *DataProcessor) AverageUtilisationLoop(
	totalTimePerCheckout *[]float64,
	avgTimePerCheckout *[]float64,
	utilisation *[]float64,
	data *CheckoutUsageData,
) {
	(*totalTimePerCheckout)[data.CheckoutNum] += data.TimeSpent
	(*avgTimePerCheckout)[data.CheckoutNum] =
		(*totalTimePerCheckout)[data.CheckoutNum] /
			(float64)(data.TotalCustomersProcessed)

	processor.computeUtilisation(totalTimePerCheckout, utilisation)
}

func (processor *DataProcessor) computeUtilisation(
	totalTimePerCheckout *[]float64, utilisation *[]float64,
) {
	for i := range *utilisation {
		(*utilisation)[i] =
			(*totalTimePerCheckout)[i] /
				utils.Sum(totalTimePerCheckout)
	}
}

// CheckoutUsageData contains data that will be used to calculate utilisation
// avergaes.
type CheckoutUsageData struct {
	CheckoutNum             int
	TimeSpent               float64
	TotalCustomersProcessed int
}
