package constants

type TransactionState uint

const (
	Pending  TransactionState = 0
	Approved TransactionState = 1
	Rejected TransactionState = 2
)
