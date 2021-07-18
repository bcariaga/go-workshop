package models

type Passenger interface {
	GetDiscount() float32
}

type BasePassenger struct {
}
type LastMinutePassenger struct {
}
type EmployeePassenger struct {
}

type EmployeeLastMinutePassenger struct {
	LastMinutePassenger
	EmployeePassenger
}

func (e BasePassenger) GetDiscount() float32 {
	return 0
}
func (e LastMinutePassenger) GetDiscount() float32 {
	return .5
}
func (e EmployeePassenger) GetDiscount() float32 {
	return 1
}
func (e EmployeeLastMinutePassenger) GetDiscount() float32 {
	return e.EmployeePassenger.GetDiscount() + e.LastMinutePassenger.GetDiscount()
}
