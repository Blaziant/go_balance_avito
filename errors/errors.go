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
