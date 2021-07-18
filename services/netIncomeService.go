package services

import "bcariaga/airlines/models"

func CalculateNetIncome(baseTicket models.Ticket, passengers []models.Passenger) float32 {
	var netIncome float32
	for _, passenger := range passengers {
		netIncome += baseTicket.CalculatePriceFor(passenger)
	}
	return netIncome
}
