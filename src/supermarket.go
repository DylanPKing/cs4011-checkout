// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

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
		CustomerData:  make(chan *agents.CustomerData),
		DataLogger:    &logger,
	}
	go dataProcessor.ComputeAverageUtilisation(storeManager.InitialNumberOfCheckouts)
	go dataProcessor.ProcessCustomerData()
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
		go func(i int, itemLimit int) {
			checkouts[i] = *agents.NewCheckout(
				itemLimit, storeManager.QueueLimit, i, &dataProcessor,
			)
			checkouts[i].ServeCustomer()
		}(i, itemLimit)
	}

	go func() {
		for {
			customer := agents.NewCustomer(&seed, &dataProcessor, weatherAgent)
			go customer.QueueCheckout(&checkouts)
			time.Sleep(
				((time.Second) * time.Duration(weatherAgent.CustomerEntryRate)) / 720,
			)
		}
	}()

	startTime := time.Now()
	endTime := startTime.Add(time.Minute * 1)

	go func() {
		for {
			time.Sleep(time.Second * time.Duration((rand.Int()*5)+10))
			weatherAgent.ToggleWeather()
		}
	}()

	for {
		now := time.Now()
		if now.After(endTime) {
			break
		}
	}

	logger.WriteOutputToFile()
}
