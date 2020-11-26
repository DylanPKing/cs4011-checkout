package agents

import (
	"sync"
	"sync/atomic"
	//"time"
)

// Checkout struct with potential item limit, 5/10 and queue limit, usually 6, possibly 4 for COVID
type Checkout struct {
	ItemLimit       int
	QueueLimit      int
	Queue           chan *Customer
	CurrentQueueLen int64
	QueueLock       sync.Mutex
}

// NewCheckout creates a checkout
func NewCheckout(itemLim int, queueLim int) *Checkout {
	var theQueue = make(chan *Customer)
	checkout := Checkout{
		ItemLimit:  itemLim,
		QueueLimit: queueLim,
		Queue:      theQueue,
	}
	return &checkout
}

// ServeCustomer serves a single customer
// TODO: Take in first customer
// TODO: Loop through the item array
// TODO: Run ScanItem method
// TODO: Shift all customers up 1 index in array/space in queue
// TODO: Pass info to Dylan's data processor
func (checkout *Checkout) ServeCustomer() {

	for {
		// BEGINNING TIME
		customer, ok := <-checkout.Queue
		if ok {
			atomic.AddInt64(&checkout.CurrentQueueLen, -1)
			for _, v := range customer.NumberOfItems {
				ScanItem(v)
			}
			// If we want to add sleep for bagging and paying, add here
			customer.checkedOut = true
		}
		// END TIME
		// SUBTRACT BEGINNING FROM END
	}
}

// ScanItem scans the current item for the current customer
// TODO: Multiply a constant by weight
func ScanItem(item *Product) {
	//time.Sleep()
}
