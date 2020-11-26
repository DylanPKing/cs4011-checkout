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
	OutputBuffer strings.Builder
	OutputFile   string
}

// LogCheckoutUtlisation logs the data produced by ComputeAverageUtilisation.
func (logger *Logger) LogCheckoutUtlisation(
	totalUtilisationPerCheckout []float64, avgTimePerCheckout []float64,
	utilisation []float64,
) {
	var output strings.Builder
	output.WriteString(fmt.Sprintln("Total utilisation of each checkout:"))
	for i, total := range totalUtilisationPerCheckout {
		output.WriteString(fmt.Sprintf("\tCheckout %d: %f\n", i, total))
	}

	output.WriteString(fmt.Sprintln("Average time spent at each checkout:"))
	for i, avg := range avgTimePerCheckout {
		output.WriteString(fmt.Sprintf("\tCheckout %d: %f\n", i, avg))
	}

	output.WriteString(fmt.Sprintln("Percent utilisation of each checkout:"))
	for i, usage := range utilisation {
		output.WriteString(fmt.Sprintf("\tCheckout %d: %f\n", i, usage))
	}

	fmt.Print(output.String())
	logger.OutputBuffer.WriteString(output.String())
}

// WriteOutputToFile will dump all of the logged data during execution to an
// output file at the end of nthe programs runtime.
func (logger *Logger) WriteOutputToFile() {
	logger.createOutputFileIfNotExists()
	bytes := []byte(logger.OutputBuffer.String())
	err := ioutil.WriteFile(logger.OutputFile, bytes, 0644)
	utils.CheckIsErrorRaised(err)
	logger.OutputBuffer.Reset()
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
		fmt.Sprintln("A customer has lost their patience and left the store"),
	)
	output.WriteString(
		fmt.Sprintf("\tTotal customers lost: %d\n", totalLostCustomers),
	)
	fmt.Print(output.String())
	logger.OutputBuffer.WriteString(output.String())
}

// LogWeatherChange will log what conditios have changed with the weather.
func (logger *Logger) LogWeatherChange(
	currentCondition string, patienceMultiplier float32,
	entryRate float32, timesChanged int,
) {
	var output strings.Builder

	output.WriteString(
		fmt.Sprintf("The weather has changed to %s.\n", currentCondition),
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
	logger.OutputBuffer.WriteString(output.String())
}
