/*
1. In go we only have for loop with different types

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Package level variables cannot be able to create with := syntax it must use var syntax
var conferenceName = "GoLang Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// Below is the code to initialize list of maps by passing any integer initial value
//var bookings = make([]map[string]string, 0)

//Below is the code to initilaize list of structs

var bookings = make([]UserData, 0)

//Struct is used to define mixed datatypes, we define strucutre of

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//We can declare variable in both ways in go by using above or below syntax
	//var conferenceName = "GoLang Conference"

	//we use %T to print type of variable
	greetUsers()
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)

	// fmt.Println(conferenceName)

	// var bookings [50]string. This is how we will define an array of fixed size
	// var bookings []string //Defining a list with dynamic size

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		/*
			In Go, the main function does not wait for other goroutines (lightweight threads) to complete.
			If the main goroutine finishes execution,
			the program exits immediately — even if other goroutines are still running.
		*/
		//You're telling Go to run sendTicket asynchronously and independently in a separate lightweight thread, called a goroutine.
		//We are specifying number of thereads/routines to add to wait group before exiting main function
		//When main thread has 0 thereads to execute then main fucntion will exit
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// bookings[0] = firstName + " " + lastName
		// fmt.Printf("The whole slice is:%v\n", bookings)
		// fmt.Printf("First element in slice is%v\n", bookings[0])
		// fmt.Printf("Type of slice is %T\n", bookings)
		// fmt.Printf("Length of an slice is %v\n", len(bookings))

		//first names logic
		firstNames := getFirstNames()
		fmt.Printf("First names of all the bookings are%v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("Our all conference tickets got booked")
			//break
		}
	} else {
		// fmt.Printf("We only have %v tickets left so you cannot book %v tickets\n", remainingTickets, userTickets)
		if !isValidEmail {
			fmt.Println("Your input email doesnot contain @ sign")
		}
		if !isValidName {
			fmt.Println("Entered first name or last name is too short")
		}
		if !isValidTicketNumber {
			fmt.Println("User has enterd invalid number of tickets")
		}

		// fmt.Println("Your inout data is invalid try again later")

	}
	//Here we are saying go to wait for all theread to complete executing that we added in wg.add to complete before exiting main theread
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v conference booking application\n", conferenceName)
	fmt.Println("We have totoal of %v tickets and %v are still available", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//Defining a map which is key-value pair which is similar to dic in python
	//In go we cannot mix datatypes under one map
	//Below is the code to initlaize user data of map type
	//var userData = make(map[string]string)
	//Below is the code to initialize userData of map type
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	//Below is the code to store details of map type
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for conference %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("######")
	fmt.Printf("Sending:\n %v ticket to email address %v\n", ticket, email)
	fmt.Printf("#######")
	//We done function removes the theraad form wg.Add saying like my executing part is Done executing.
	wg.Done()
}
