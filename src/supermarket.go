package main

import (
	"./agents"
	"./manager"

	"math/rand"
	"time"
)

func main() {
	// Seed the random number
	seed := rand.NewSource(time.Now().UnixNano())

	const loggerOutput = "../out/loggeroutput"

	storeManager := manager.NewManager()
	storeManager.StartCheckouts()

	logger := agents.Logger{
		OutputFile: loggerOutput,
	}

	dataProcessor := agents.DataProcessor{
		CheckoutUsage: make(chan *agents.CheckoutUsageData),
		// CustomerData: make(chan *agents.Customer),
		DataLogger: &logger,
	}

	weatherAgent := agents.NewWeather(&seed, &dataProcessor)
	weatherAgent.ToggleWeather()
	// Weather agent legend
	// You can get the following information from the weather agent:
	// TimesChangedToday          int
	// CustomerPatienceMultiplier float32
	// CustomerEntryRate          float32
	// Conditions                 *Condition
	// CurrentCondition           string
	// Seed                       *rand.Source
	// NOTE: The purpose of Seed in the model is to eliminate some of the passing around of the seed value as a paramater - speciffically the ToggleWeather function

	logger.WriteOutputToFile()
}
