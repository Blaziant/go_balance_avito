package errors

type NotEnoughtMoneyError struct {
	Message string
}

func (e *NotEnoughtMoneyError) Error() string {
	return e.Message
}

type NegativePrice struct {
	Message string
}

func (e *NegativePrice) Error() string {
	return e.Message
}

type AcceptedOrder struct {
	Message string
}

func (e *AcceptedOrder) Error() string {
	return e.Message
}

type ReserveAccountUsed struct {
	Message string
}

func (e *ReserveAccountUsed) Error() string {
	return e.Message
}
