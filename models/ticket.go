package models

type Ticket struct {
	BasePrice float32
}

func (t Ticket) CalculatePriceFor(passenger Passenger) float32 {
	return t.BasePrice - (t.BasePrice * passenger.GetDiscount())
}
