package main

import (
	"fmt"
	"math/rand"
	"time"

	"./agents"
	"./manager"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	storeManager := manager.NewManager()
	storeManager.StartCheckouts()

	testCustomer := agents.NewCustomer(&seed)
	fmt.Println(testCustomer.Patience)
	fmt.Println(len(testCustomer.NumberOfItems))
	fmt.Println(testCustomer.NumberOfItems[0].Weight)
}
