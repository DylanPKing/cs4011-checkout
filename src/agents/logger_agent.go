package agents

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../utils"
)

// Logger is used to print data to the console and to an output file.
type Logger struct {
	UtilisationBuffer   strings.Builder
	LostCustomerBuffer  strings.Builder
	WeatherChangeBuffer strings.Builder
	CustomerDataBuffer  strings.Builder
	OutputFile          string
}

// LogCheckoutUtlisation logs the data produced by ComputeAverageUtilisation.
func (logger *Logger) LogCheckoutUtlisation(
	totalUtilisationPerCheckout []float64, avgTimePerCheckout []float64,
	utilisation []float64,
) {
	var output strings.Builder
	output.WriteString(fmt.Sprintln("\n\nTotal time spent at each checkout:"))
	for i, total := range totalUtilisationPerCheckout {
		output.WriteString(fmt.Sprintf("\tCheckout %d: %.2fs\n", i, total))
	}

	output.WriteString(fmt.Sprintln("Average time spent at each checkout:"))
	for i, avg := range avgTimePerCheckout {
		output.WriteString(fmt.Sprintf("\tCheckout %d: %.2fs\n", i, avg))
	}

	output.WriteString(fmt.Sprintln("Percent utilisation of each checkout:"))
	for i, usage := range utilisation {
		output.WriteString(
			fmt.Sprintf("\tCheckout %d: %.2f%%\n", i, usage*100),
		)
	}

	fmt.Print(output.String())
	logger.UtilisationBuffer.Reset()
	logger.UtilisationBuffer.WriteString(output.String())
}

// WriteOutputToFile will dump all of the logged data during execution to an
// output file at the end of nthe programs runtime.
func (logger *Logger) WriteOutputToFile() {
	logger.createOutputFileIfNotExists()

	var megaBuffer strings.Builder

	megaBuffer.WriteString(logger.UtilisationBuffer.String())
	megaBuffer.WriteString(logger.CustomerDataBuffer.String())
	megaBuffer.WriteString(logger.LostCustomerBuffer.String())
	megaBuffer.WriteString(logger.WeatherChangeBuffer.String())

	bytes := []byte(megaBuffer.String())
	err := ioutil.WriteFile(logger.OutputFile, bytes, 0644)
	utils.CheckIsErrorRaised(err)
}

func (logger *Logger) createOutputFileIfNotExists() {
	file, err := os.Create(logger.OutputFile)
	utils.CheckIsErrorRaised(err)
	file.Close()
}

// LogCustomerLost Logs when a customer leaves the store without buying
// anything, and the total number of customers lost.
func (logger *Logger) LogCustomerLost(totalLostCustomers int64) {
	var output strings.Builder
	output.WriteString(
		fmt.Sprintf("\n\nTotal customers lost: %d\n", totalLostCustomers),
	)
	fmt.Print(output.String())
	logger.LostCustomerBuffer.Reset()
	logger.LostCustomerBuffer.WriteString(output.String())
}

// LogWeatherChange will log what conditios have changed with the weather.
func (logger *Logger) LogWeatherChange(
	currentCondition string, patienceMultiplier float32,
	entryRate float32, timesChanged int,
) {
	var output strings.Builder

	output.WriteString(
		fmt.Sprintf("\n\nThe weather has changed to %s.\n", currentCondition),
	)
	output.WriteString(
		fmt.Sprintf(
			"\tNew customer patience multipler: %.2f\n", patienceMultiplier,
		),
	)
	output.WriteString(
		fmt.Sprintf("\tNew customer entry rate: %.2f\n", entryRate),
	)
	output.WriteString(
		fmt.Sprintf("\tTotal times weather has changed: %d\n", timesChanged),
	)
	fmt.Print(output.String())
	logger.WeatherChangeBuffer.Reset()
	logger.WeatherChangeBuffer.WriteString(output.String())
}

// LogCustomerData Logs the totals and averages calculated by the
// data processor
func (logger *Logger) LogCustomerData(
	totalProductsProcessed int, averageProductsPerTrolley int,
	averageWaitTime float64, waitTimeforCustomer float64,
	totalCustomers int,
) {
	var output strings.Builder

	output.WriteString(
		fmt.Sprintf(
			"\n\nWait time for current customer: %.2fs\n", waitTimeforCustomer,
		),
	)
	output.WriteString(
		fmt.Sprintf("Total products processed: %d\n", totalProductsProcessed),
	)
	output.WriteString(
		fmt.Sprintf(
			"Average products per trolley: %d\n", averageProductsPerTrolley,
		),
	)
	output.WriteString(
		fmt.Sprintf(
			"Average wait time for a customer: %.2fs\n", averageWaitTime,
		),
	)
	output.WriteString(
		fmt.Sprintf("Total customers processed: %d\n", totalCustomers),
	)
	fmt.Print(output.String())
	logger.CustomerDataBuffer.Reset()
	logger.CustomerDataBuffer.WriteString(output.String())
}
