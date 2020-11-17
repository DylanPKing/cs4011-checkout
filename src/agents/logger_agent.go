package agents

import (
	"fmt"
	"strings"
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
