package main

import (
	"os"

	"./agents"
	"./manager"

	"math/rand"
	"time"
)

func main() {
	// Seed the random number
	seed := rand.NewSource(time.Now().UnixNano())
	workingDir, _ := os.Getwd()
	loggerOutput := (workingDir + "\\out\\loggeroutput")

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
	go dataProcessor.ComputeAverageUtilisation()
	weatherAgent := agents.NewWeather(&seed, &dataProcessor)
	weatherAgent.ToggleWeather()

	// Set up checkouts
	checkouts := make([]agents.Checkout, storeManager.InitialNumberOfCheckouts)
	var itemLimit int
	for i := 0; i < storeManager.InitialNumberOfCheckouts; i++ {
		if i < storeManager.NumberOfExpressCheckouts {
			itemLimit = storeManager.NumberOfExpressItems
		} else {
			itemLimit = 200
		}
		go func(i int) {
			checkouts[i] = *agents.NewCheckout(itemLimit, storeManager.QueueLimit)
			checkouts[i].ServeCustomer()
		}(i)
	}

	go addCustomer(&seed, &checkouts)

	startTime := time.Now()
	endTime := startTime.Add(time.Minute * 1)
	for {
		now := time.Now()
		if now.After(endTime) {
			break
		}
	}

	logger.WriteOutputToFile()
}

func addCustomer(seed *rand.Source, checkouts *[]agents.Checkout) {
	for {
		customer := agents.NewCustomer(seed)
		customer.QueueCheckout(checkouts)
	}

}
