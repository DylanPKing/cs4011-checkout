package test

import (
	"io/ioutil"
	"testing"

	"../../src/agents"
	"../../src/utils"
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

func Test_WriteOutputToFile(t *testing.T) {
	filePath := "../../out/test_output"

	logger := &agents.Logger{
		OutputFile: filePath,
	}

	expectedOutput := "Test output"

	logger.OutputBuffer.WriteString(expectedOutput)

	logger.WriteOutputToFile()

	fileBytes, err := ioutil.ReadFile(filePath)
	utils.CheckIsErrorRaised(err)

	actualOutput := string(fileBytes)

	assert.Equal(t, expectedOutput, actualOutput)
}

func Test_LogCustomerLost(t *testing.T) {
	logger := agents.Logger{}

	logger.LogCustomerLost(int64(1))

	expectedOutput := "A customer has lost their patience and left the " +
		"store\n\tTotal customers lost: 1\n"

	actualOutput := logger.OutputBuffer.String()

	assert.Equal(t, actualOutput, expectedOutput)
}
