// Till No. 7
// Dylan King - 17197813
// Louise Madden - 17198232
// Brian Malone - 17198178
// Szymon Sztyrmer - 17200296

package agents

import (
	"math/rand"
	"sync/atomic"
	"time"
)

//Customer defines a customer
type Customer struct {
	Patience      float32
	NumberOfItems []*Product
	Queue         bool
	checkedOut    bool
	dataProcessor *DataProcessor
	TotalWaitTime float32
	WeatherAgent  *Weather
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
func NewCustomer(seed *rand.Source, processor *DataProcessor, weather *Weather) *Customer {
	p := ((rand.Float32() * 10) + 3) * weather.CustomerPatienceMultiplier
	t := FillTrolley(seed)
	customer := Customer{
		Patience:      p,
		NumberOfItems: t,
		Queue:         false,
		dataProcessor: processor,
		WeatherAgent:  weather,
	}
	return &customer
}

//NewProduct returns a product
func NewProduct(seed *rand.Source) *Product {
	w := (rand.Float32() * 12) + 1.0
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
func (customer *Customer) ToggleQueue() {
	customer.Queue = !(customer.Queue)
}

//QueueCheckout is for queueing at a checkout
func (customer *Customer) QueueCheckout(checkouts *[]Checkout) {
	shortestQueueLength := (*checkouts)[0].CurrentQueueLen
	indexCheckout := 0
	for {
		for j := 1; j < len(*checkouts); j++ {
			if (*checkouts)[j].CurrentQueueLen < shortestQueueLength &&
				len(customer.NumberOfItems) <= (*checkouts)[j].ItemLimit {
				shortestQueueLength = (*checkouts)[j].CurrentQueueLen
				indexCheckout = j
			}
		}
		if int(shortestQueueLength) < (*checkouts)[indexCheckout].QueueLimit {
			(*checkouts)[indexCheckout].Queue <- customer
			atomic.AddInt64(&(*checkouts)[indexCheckout].CurrentQueueLen, 1)
			break
		} else {
			customer.Patience -= (2 * customer.WeatherAgent.CustomerPatienceMultiplier)
			time.Sleep(time.Duration(time.Second / 720))
			if customer.Patience <= 0 {
				leaveStore(customer, false)
				return
			}
		}
	}
	startTime := time.Now().Unix()
	for {
		customer.Patience -= (2 * customer.WeatherAgent.CustomerPatienceMultiplier)
		time.Sleep(
			(time.Second * (time.Duration(len(customer.NumberOfItems))) * 2) / 720,
		)
		if customer.checkedOut {
			endTime := time.Now().Unix()
			customer.TotalWaitTime = float32(endTime - startTime)
			leaveStore(customer, true)
			break
		} else if customer.Patience <= 0 {
			endTime := time.Now().Unix()
			customer.TotalWaitTime = float32(endTime - startTime)
			leaveStore(customer, false)
			break
		}
	}
}

func leaveStore(customer *Customer, happy bool) {
	if happy {
		customerData := CustomerData{
			NumberOfItems: len(customer.NumberOfItems),
			TotalWaitTime: float64(customer.TotalWaitTime),
		}
		customer.dataProcessor.CustomerData <- &customerData
	} else {
		go customer.dataProcessor.IncrementLostCustomers()
	}
}
