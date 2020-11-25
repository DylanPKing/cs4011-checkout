package agents

import "math/rand"

//Customer defines a customer
type Customer struct {
	Patience      float32
	NumberOfItems []*Product
	Queue         bool
	checkedOut    bool
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

//write a method for choosing a checkout
func (customer *Customer) QueueCheckout(checkout *Checkout) {
	shortestQueue := CurrentNumberOfCheckouts[0]
	indexCheckout := 0
	for j = 1; j < CurrentNumberOfCheckouts.len(); j++ {
		if CurrentNumberOfCheckouts[j].Checkout.CurrentQueueLen < shortestQueue {
			shortestQueue = CurrentNumberOfCheckouts[j]
			indexCheckout = j
		}
	}
	checkout.JoinCheckout(customer)
	for {
		customer.Patience--
		//time.sleep()
		if customer.checkedOut || customer.Patience == 0 {
			leaveStore(customer)
		}

	}

}

func leaveStore(customer *Customer) {

}
