package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to our flight booking application")
	fmt.Println("Book your joy rides here")

	var command = "go run <file name ex. main.go>"
	fmt.Println("to run the go file run this command : ", command)

	const myNumber = 7

  // declaring a string variable in the shortest form
	bookingClass := "FIRST CLASS" // this can be also written in verbose like this - var bookingClass string = "FIRST CLASS"

	// myNumber = 28 cannot resign a variable declared with const
	fmt.Println(myNumber)

	// const value can't be reassigned
	const serviceName = "GO Beach Rides"

	var totalSeats uint = 50
	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var availableSeats uint

		// var bookings [50]string

		// Slices don't require declaring the size

		fmt.Print("\nwhat's your first name?\n")
		fmt.Scan(&firstName)

		fmt.Print("\nwhat's your last name?\n")
		fmt.Scan(&lastName)

		fmt.Print("\nwhat's your email?\n")
		fmt.Scan(&email)

		fmt.Print("\nHow many tickets you'd like to book today?\n")
		fmt.Scan(&availableSeats)

		totalSeats = totalSeats - availableSeats
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("\nThank you %v for your booking. You will recieve the booking confirmation to your email : %v \n", firstName, email)
		fmt.Printf("\nHello %v. Your booking details are below. \nThank you for Choosing us.\n", firstName)

		fmt.Printf("\nBooking Class : %v \nTotal tickets booked : %v \n", bookingClass, availableSeats)
		fmt.Printf("\nTotal Available Seats: %v\n", totalSeats)

		// fmt.Printf("Printing whole slice of bookings: %v \n", bookings)
		// fmt.Printf("First Booking: %v\n", bookings[0])
		// fmt.Printf("Slice Type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		// Go warns you even before executing this program about the Index of 52 is out of bounds

		//  fmt.Printf("Printing the index gt. then the actual fixed array value: %v", bookings[52])

		// Slices are more convineient and dynamic and we don't need to use the index of the element

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all the bookings : %v\n", firstNames)

	}

}
