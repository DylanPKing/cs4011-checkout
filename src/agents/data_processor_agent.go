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
	var totalTimePerCheckout []float64
	var avgTimePerCheckout []float64
	var utilisation []float64
	for {
		data := <-processor.CheckoutUsage

		totalTimePerCheckout[data.CheckoutNum] += data.TimeSpent
		avgTimePerCheckout[data.CheckoutNum] =
			totalTimePerCheckout[data.CheckoutNum] /
				(float64)(data.TotalCustomersProcessed)

		utilisation[data.CheckoutNum] =
			totalTimePerCheckout[data.CheckoutNum] /
				utils.Sum(&totalTimePerCheckout)
	}
}

// CheckoutUsageData contains data that will be used to calculate utilisation
// avergaes.
type CheckoutUsageData struct {
	CheckoutNum             int
	TimeSpent               float64
	TotalCustomersProcessed int
}
