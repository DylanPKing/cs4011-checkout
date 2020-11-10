package 

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

// NewCustomer creates a new customer
func NewCustomer(seed *rand.Source) *Customer {
	p := generatePatience(seed)
	t := fillTrolley(seed)
	customer := Customer{
		Patience:      p,
		NumberOfItems: t,
		Queue:         false,
	}
	return &customer
}

//generatePatience randomly generates a customers patience
func generatePatience(seed *rand.Source) float32 {
	randomNum := (rand.New(*seed)).Intn(10)
	return float32(randomNum)
}

//newProduct returns a product
func newProduct(seed *rand.Source) *Product {
	w := generateWeight(seed)
	product := Product{
		Weight: w,
	}
	return &product
}

//generateWeight randomly generates a products weight
func generateWeight(seed *rand.Source) float32 {
	randomNum := (rand.New(*seed)).Intn(12)
	return float32(randomNum)
}

//generateNumItems randomly generates the number of items
func generateNumItems(seed *rand.Source) int {
	randomNum := (rand.New(*seed)).Intn(200) + 1
	return randomNum
}

//fillTrolley fills a customers trolley with a random number of items of a random weight
func fillTrolley(seed *rand.Source) []*Product {
	NumberOfItems := generateNumItems(seed)
	trolley := make([]*Product, NumberOfItems)
	for i := 0; i < NumberOfItems; i++ {
		trolley[i] = newProduct(seed)
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