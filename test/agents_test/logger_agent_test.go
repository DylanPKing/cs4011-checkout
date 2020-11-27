// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

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

	expectedOutput := "\n\nTotal time spent at each checkout:\n" +
		"\tCheckout 0: 10.00s\n" +
		"\tCheckout 1: 10.00s\n" +
		"\tCheckout 2: 10.00s\n" +
		"\tCheckout 3: 10.00s\n" +
		"Average time spent at each checkout:\n" +
		"\tCheckout 0: 2.00s\n" +
		"\tCheckout 1: 2.00s\n" +
		"\tCheckout 2: 2.00s\n" +
		"\tCheckout 3: 2.00s\n" +
		"Percent utilisation of each checkout:\n" +
		"\tCheckout 0: 25.00%\n" +
		"\tCheckout 1: 25.00%\n" +
		"\tCheckout 2: 25.00%\n" +
		"\tCheckout 3: 25.00%\n"

	logger.LogCheckoutUtlisation(total, avg, utilisation)

	actualOutput := logger.UtilisationBuffer.String()

	assert.Equal(t, actualOutput, expectedOutput)
}

func Test_WriteOutputToFile(t *testing.T) {
	filePath := "../../out/test_output"

	logger := &agents.Logger{
		OutputFile: filePath,
	}

	expectedOutput := "Test output"

	logger.UtilisationBuffer.WriteString(expectedOutput)

	logger.WriteOutputToFile()

	fileBytes, err := ioutil.ReadFile(filePath)
	utils.CheckIsErrorRaised(err)

	actualOutput := string(fileBytes)

	assert.Equal(t, expectedOutput, actualOutput)
}

func Test_LogCustomerLost(t *testing.T) {
	logger := agents.Logger{}

	logger.LogCustomerLost(int64(1))

	expectedOutput := "\n\nTotal customers lost: 1\n"

	actualOutput := logger.LostCustomerBuffer.String()

	assert.Equal(t, actualOutput, expectedOutput)
}

func Test_LogWeatherChange(t *testing.T) {
	logger := agents.Logger{}

	logger.LogWeatherChange("TestCondition", 1.0, 1.0, 1)

	expectedOutput := "\n\nThe weather has changed to TestCondition.\n" +
		"\tNew customer patience multipler: 1.00\n" +
		"\tNew customer entry rate: 1.00\n" +
		"\tTotal times weather has changed: 1\n"

	actualOutput := logger.WeatherChangeBuffer.String()

	assert.Equal(t, actualOutput, expectedOutput)
}
