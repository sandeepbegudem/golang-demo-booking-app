package main

import (
	"fmt"
	"strings"
)

func main() {

	bookingClass := "PREMIUM CLASS"

	var greeting string = "Hi there!.. Greeting of the day!\n"
	fmt.Print(greeting)

	const serviceName = "Amzing Rides"

	var totalSeats uint = 50
	bookings := []string{}

	greetUsers()

	for {
		var firstName string
		var lastName string
		var email string
		var availableSeats uint

		fmt.Print("\nwhat's your first name?\n")
		fmt.Scan(&firstName)

		fmt.Print("\nwhat's your last name?\n")
		fmt.Scan(&lastName)

		fmt.Print("\nwhat's your email?\n")
		fmt.Scan(&email)

		fmt.Print("\nHow many tickets you'd like to book today?\n")
		fmt.Scan(&availableSeats)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@mail.com")
		isValidTicketEntry := availableSeats > 0 && availableSeats <= totalSeats

		if isValidName && isValidEmail && isValidTicketEntry {
			totalSeats = totalSeats - availableSeats
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v for your booking. You will recieve the booking confirmation to your email : %v \n", firstName, email)
			fmt.Printf("\nHello %v. Your booking details are below. \nThank you for Choosing us.\n", firstName)

			fmt.Printf("\nBooking Class : %v \nTotal tickets booked : %v \n", bookingClass, availableSeats)
			fmt.Printf("\nTotal Available Seats: %v\n", totalSeats)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all the bookings : %v\n", firstNames)

			if totalSeats == 0 {
				fmt.Print("We regret to inform you there are no more bookings available, please come back next time!\n")
				fmt.Print("We appreciate your patience!\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("your first name and or last name is too short. Ex: john")
			}
			if !isValidEmail {
				fmt.Println("email address you entered is invalid. Ex: doe@mail.com")
			}
			if !isValidTicketEntry {
				fmt.Println("enter a valid number of tickets. Ex: 5")
			}
			fmt.Printf("your input data is invalid, please check and try again later!")
			continue
		}

	}

}

func greetUsers() {
	fmt.Println("Hello, user welcome to our rides ticket booking system.")
}
