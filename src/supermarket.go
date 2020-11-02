package main

import (
	"./manager"
)

func main() {
	storeManager := manager.NewManager()
	storeManager.StartCheckouts()
}
