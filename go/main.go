main

import (
	"fmt"
	"github.com/NameSurname/Assignment1/Hotel"
	"github.com/NameSurname/Assignment1/Employee"
	"github.com/NameSurname/Assignment1/Gym"
	"github.com/NameSurname/Assignment1/Wallet"
)

func main() {
	fmt.Println("=== Advanced Programming 1 - Assignment 1 ===")
	fmt.Println("Select a task to run:")
	fmt.Println("1. Hotel Room Reservation System")
	fmt.Println("2. Employee Salary Calculator")
	fmt.Println("3. Gym Membership Management")
	fmt.Println("4. Wallet Transaction Simulation")
	fmt.Println("5. Exit")
	
	var choice int
	fmt.Print("Enter your choice (1-5): ")
	fmt.Scan(&choice)
	
	switch choice {
	case 1:
		Hotel.RunHotelSystem()
	case 2:
		Employee.RunEmployeeSystem()
	case 3:
		Gym.RunGymSystem()
	case 4:
		Wallet.RunWalletSystem()
	case 5:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid choice!")
	}
}