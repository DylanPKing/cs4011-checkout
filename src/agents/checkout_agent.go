package agents

// Checkout struct with potential item limit, 5/10 and queue limit, usually 6, possibly 4 for COVID
type Checkout struct {
	ItemLimit  int
	QueueLimit int
	Queue [QueueLimit]Customer
}

// TODO: create an array of customers with a buffer of 6/ variable that is passed in? COVID
// TODO: be able to scan items, time/speed, multiplied by the item size multiplier
// TODO: check the time it took for each customer

// NewCheckout creates a checkout
func NewCheckout(itemLim int, queueLim int) *Checkout {
	var theQueue [QueueLimit]Customer
	checkout := Checkout{
		ItemLimit:  itemLim,
		QueueLimit: queueLim,
		Queue : theQueue,
	}
	return &checkout
}

// JoinCheckout takes in customer
// TODO: add to queue array, loop and find first nil??? assign to that index
func (checkout *Checkout) JoinCheckout(customer *Customer) {
	for i = 0; i < checkout.Queue.length; i++ {
		if checkout.Queue[i] == nil {
			checkout[i] = customer
			break;
		}
	}
}

// ServeCustomer serves a single customer
// TODO: Take in first customer
// TODO: Loop through the item array
// TODO: Run ScanItem method
// TODO: Shift all customers up 1 index in array/space in queue
// TODO: Pass info to Dylan's data processor
func (checkout *Checkout) ServeCustomer() {
	// BEGINNING TIME
	current := Queue[0] //current is customer
	for i, v := range current.Trolley {
		ScanItem(v)
	}
	// END TIME
	// SUBTRACT BEGINNING FROM END

	// If we want to add sleep for bagging and paying, add here
	checkout.MoveQueue()
	LeaveStore(current)
}

// ScanItem scans the current item for the current customer
// TODO: Multiply a constant by weight
func ScanItem(item *Product) {

}

// MoveQueue removes the first customer(index 0) from the queue
// TODO: Remove first customer
// TODO: Move other customers up an index
func (checkout *Checkout) MoveQueue() {
	newQueue := [QueueLimit]Customer

	for i = 1; i < checkout.Queue.length; i++ {
		newQueue[i-1] = chechout.Queue[i]
		checkout.Queue = newQueue
	}
}

// LeaveStore is the customer leaving the store, and calling data processor
// TODO: Have an if statement to see if customer went through checkout or left angry
// TODO: If through checkout send what is necessary for Dylan
// TODO: If angry send that they were angry
func LeaveStore(customer *Customer) {

}
