package test

import (
	"testing"

	"../../src/agents"
	"github.com/stretchr/testify/assert"
)

func Test_LogCheckoutUtilisation(t *testing.T) {
	logger := agents.Logger{
		OutputFile: "out.txt",
	}

	total := make([]float64, 4)
	avg := make([]float64, 4)
	utilisation := make([]float64, 4)

	for i := range total {
		total[i] = 10.0
		avg[i] = 2.0
		utilisation[i] = 0.25
	}

	expectedOutput := "Total utilisation of each checkout:\n" +
		"\tCheckout 0: 10.000000\n" +
		"\tCheckout 1: 10.000000\n" +
		"\tCheckout 2: 10.000000\n" +
		"\tCheckout 3: 10.000000\n" +
		"Average time spent at each checkout:\n" +
		"\tCheckout 0: 2.000000\n" +
		"\tCheckout 1: 2.000000\n" +
		"\tCheckout 2: 2.000000\n" +
		"\tCheckout 3: 2.000000\n" +
		"Percent utilisation of each checkout:\n" +
		"\tCheckout 0: 0.250000\n" +
		"\tCheckout 1: 0.250000\n" +
		"\tCheckout 2: 0.250000\n" +
		"\tCheckout 3: 0.250000\n"

	logger.LogCheckoutUtlisation(total, avg, utilisation)

	actualOutput := logger.OutputBuffer.String()

	assert.Equal(t, actualOutput, expectedOutput)
}
