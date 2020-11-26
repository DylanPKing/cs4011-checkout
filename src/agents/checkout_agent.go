package agents

import (
	"sync/atomic"
	"time"
)

// Checkout struct with potential item limit, 5/10 and queue limit, usually 6, possibly 4 for COVID
type Checkout struct {
	Number                  int
	ItemLimit               int
	QueueLimit              int
	Queue                   chan *Customer
	CurrentQueueLen         int64
	TotalCustomersProcessed int
	processor               *DataProcessor
}

// NewCheckout creates a checkout
func NewCheckout(itemLim int, queueLim int, number int, dataProcessor *DataProcessor) *Checkout {
	var theQueue = make(chan *Customer)
	checkout := Checkout{
		Number:     number,
		ItemLimit:  itemLim,
		QueueLimit: queueLim,
		Queue:      theQueue,
		processor:  dataProcessor,
	}
	return &checkout
}

// ServeCustomer serves a single customer
func (checkout *Checkout) ServeCustomer() {

	for {
		customer, ok := <-checkout.Queue
		if ok {
			startTime := time.Now().Unix()
			for _, v := range customer.NumberOfItems {
				ScanItem(v)
			}
			atomic.AddInt64(&checkout.CurrentQueueLen, -1)

			customer.checkedOut = true
			checkout.TotalCustomersProcessed++
			endTime := time.Now().Unix()
			data := &CheckoutUsageData{
				CheckoutNum:             checkout.Number,
				TimeSpent:               float64(endTime - startTime),
				TotalCustomersProcessed: checkout.TotalCustomersProcessed,
			}
			checkout.processor.CheckoutUsage <- data
		}
	}
}

// ScanItem scans the current item for the current customer
func ScanItem(item *Product) {
	timeToPause := (time.Second * time.Duration(item.Weight*0.5)) / 720
	time.Sleep(timeToPause)
}
