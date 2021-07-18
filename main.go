package main

import (
	"bcariaga/airlines/services"
	"fmt"
)

func main() {
	var basePrice float32 = 1000
	fmt.Println("calculating net income...")
	fmt.Printf("base price: %v \n", basePrice)
	fmt.Printf("net income: %v \n", ProccessIncome(basePrice))
	fmt.Println("finish!")
}

func ProccessIncome(ticketPrice float32) float32 {
	return services.CalculateNetIncome(
		*services.CreateTicket(ticketPrice),
		*services.CreatePassengers())
}
