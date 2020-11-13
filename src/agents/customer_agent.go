package agents

import "math/rand"

//Customer defines a customer
type Customer struct {
	Patience      float32
	NumberOfItems []*Product
	Queue         bool
}

//Product defines a product
type Product struct {
	Weight float32
}

// RandnumGen generates a random number for given parameter
func RandnumGen(seed *rand.Source, num int) int {
	randomNum := (rand.New(*seed)).Intn(num)
	return randomNum
}

// NewCustomer creates a new customer
func NewCustomer(seed *rand.Source) *Customer {
	p := float32(RandnumGen(seed, 10))
	t := FillTrolley(seed)
	customer := Customer{
		Patience:      p,
		NumberOfItems: t,
		Queue:         false,
	}
	return &customer
}

//NewProduct returns a product
func NewProduct(seed *rand.Source) *Product {
	w := float32(RandnumGen(seed, 12) + 1)
	product := Product{
		Weight: w,
	}
	return &product
}

//FillTrolley fills a customers trolley with a random number of items of a random weight
func FillTrolley(seed *rand.Source) []*Product {
	NumberOfItems := RandnumGen(seed, 200) + 1
	trolley := make([]*Product, NumberOfItems)
	for i := 0; i < NumberOfItems; i++ {
		trolley[i] = NewProduct(seed)
	}
	return trolley
}

//ToggleQueue toggles whether customer is queueing
func (c *Customer) ToggleQueue() {
	c.Queue = !(c.Queue)
}

//Now how does one implement

//Entering the store
//Probably dealt with in the main when a new customer is called

//Going to the checkout
//when trolley is filled
//go to checkout
//compare queues, probably array of queues of open checkouts
//is added to checkout for handling

//Leaving the store
//to be decided
