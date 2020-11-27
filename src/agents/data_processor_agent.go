// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

package agents

import (
	"sync/atomic"

	"../utils"
)

// DataProcessor collects and performs calculations on data
// before sending data to the Logger
type DataProcessor struct {
	CheckoutUsage          chan *CheckoutUsageData
	CustomerData           chan *CustomerData
	AvgCheckoutUseTime     float64
	DataLogger             *Logger
	LostCustomers          int64
	totalProductsProcessed int64
}

// ComputeAverageUtilisation collects the total usage of each checkout and
// computes their average.
func (processor *DataProcessor) ComputeAverageUtilisation(numberOfCheckouts int) {
	totalTimePerCheckout := make([]float64, numberOfCheckouts)
	avgTimePerCheckout := make([]float64, numberOfCheckouts)
	utilisation := make([]float64, numberOfCheckouts)
	for {
		data, ok := <-processor.CheckoutUsage
		if ok {
			processor.AverageUtilisationLoop(
				&totalTimePerCheckout, &avgTimePerCheckout, &utilisation, data,
			)
			processor.DataLogger.LogCheckoutUtlisation(
				totalTimePerCheckout, avgTimePerCheckout, utilisation,
			)
		} else {
			break
		}
	}
}

// AverageUtilisationLoop is the core loop used to calculate the total, and
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

// IncrementLostCustomers updates the number of customers that have given up
// and left.
func (processor *DataProcessor) IncrementLostCustomers() {
	atomic.AddInt64(&processor.LostCustomers, 1)
	processor.DataLogger.LogCustomerLost(processor.LostCustomers)
}

// ProcessWeatherChange receives the details of the most recetn weather change,
// then sends it to the logger.
func (processor *DataProcessor) ProcessWeatherChange(
	currentCondition string, patienceMultiplier float32,
	entryRate float32, timesChanged int,
) {
	processor.DataLogger.LogWeatherChange(
		currentCondition, patienceMultiplier,
		entryRate, timesChanged,
	)
}

// ProcessCustomerData calculates totals and averages every time a customer is
// checked out
func (processor *DataProcessor) ProcessCustomerData() {
	totalProductsProcessed := 0
	totalCustomers := 0
	totalWaitTime := 0.0
	for {
		customerData, ok := <-processor.CustomerData
		if ok {
			totalCustomers++
			totalProductsProcessed += customerData.NumberOfItems
			averageProductsPerTrolley := totalProductsProcessed / totalCustomers
			totalWaitTime += float64(customerData.TotalWaitTime)
			averageWaitTime := totalWaitTime / float64(totalCustomers)
			processor.DataLogger.LogCustomerData(
				totalProductsProcessed, averageProductsPerTrolley,
				averageWaitTime, customerData.TotalWaitTime, totalCustomers,
			)
		} else {
			break
		}
	}
}

// CheckoutUsageData contains data that will be used to calculate utilisation
// averages.
type CheckoutUsageData struct {
	CheckoutNum             int
	TimeSpent               float64
	TotalCustomersProcessed int
}

// CustomerData contains the Data that will be used to calclate custoemr totals
// and averages
type CustomerData struct {
	NumberOfItems int
	TotalWaitTime float64
}
