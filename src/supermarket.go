package main

import (
	"fmt"

	"./manager"
)

func main() {
	fmt.Println("Hello")
	storeManager := manager.NewManager()
	storeManager.StartCheckouts()
}
