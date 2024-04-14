package main

import "fmt"

func main() {
	var son int = 69
	var son1 string = "schools"
	var conferenceName string = "go conference"
	var conferenceTicket int = 100
	const remaningTIcket = 100

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 4, 6, 7}

	son = 7

	fmt.Printf("Tonye has %v %v all over Nigeria\n", son, son1)
	fmt.Printf("Welcome to %v Booking application where we have %v booking Tickets and the remaning booking tickets are %v\n", conferenceName, conferenceTicket, remaningTIcket)

	fmt.Println(arr1)
	fmt.Println(arr2)
	var Firstname = "Noel"
	var  Lastname = "Tonye"
	fmt.Println("Hello", Firstname, Lastname, "Thank you for booking ")

	for u := 1; u <= 5; u++ {
		fmt.Println("Tboy")
	}
fmt.Println("Hello Wrold!")
}
