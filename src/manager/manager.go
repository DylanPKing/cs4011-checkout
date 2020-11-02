// Package manager controls the initial checkout configuration for the supermarket
package manager

import (
	"fmt"
	"regexp"
	"strconv"
)

// Manager is the struct that contains all the initial values that are supplied by the user
type Manager struct {
	InitialNumberOfCheckouts int
	NumberOfExpressCheckouts int
}

// NewManager creates a new manager struct
func NewManager() *Manager {
	manager := Manager{}
	return &manager
}

// StartCheckouts asks the user to initialise the number of checkouts.
func (manager *Manager) StartCheckouts() {
	for {
		manager.InitialNumberOfCheckouts = userInputInt("Please input the initial number of operating checkouts [1-8]")
		manager.NumberOfExpressCheckouts = userInputInt("Please input the number of express checkouts(no more than 5 items)")
		if manager.NumberOfExpressCheckouts > manager.InitialNumberOfCheckouts {
			fmt.Println("You can't have more express checkouts than there are total checkouts dingus...\nTry to put in some logical values ya?")
			continue
		}
		break
	}
}

// userInputInt asks the user to input an int to use as an initial value
func userInputInt(text string) int {
	regex, _ := regexp.Compile("[0-9]")
	var userInput string
	var number int

	for {
		fmt.Println(text)
		fmt.Scanln(&userInput)
		// Check for regex
		check := regex.MatchString(userInput)
		if !check {
			fmt.Println("Error: You were supposed to put in an integer number in range 1-8 dingus -_-, \nTry again...")
			continue
		}
		// Check for an integer
		inputNumber, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Error: Really mate? Do you really want a piece of a checkout? No can do chief... Input a whole number next time.")
			continue
		}
		// Check for range
		if 1 > inputNumber || inputNumber > 8 {
			fmt.Println("*sigh*...\nThe instructions literally told you what range to USE!\nWhy would you even try that?...\nBack to the top you go. Try to do it properly this time.")
			continue
		}
		number = inputNumber
		break
	}
	return number
}
