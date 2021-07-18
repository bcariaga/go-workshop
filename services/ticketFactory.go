package services

import "bcariaga/airlines/models"

func CreateTicket(basePrice float32) *models.Ticket {
	return &models.Ticket{BasePrice: basePrice}
}
