package services

import "bcariaga/airlines/models"

func CreatePassengers() *[]models.Passenger {
	return &[]models.Passenger{
		&models.BasePassenger{},
		&models.BasePassenger{},
		&models.EmployeePassenger{},
		&models.LastMinutePassenger{},
		&models.EmployeeLastMinutePassenger{},
	}
}
