package main

import "fmt"

func main() {
  fmt.Println("Welcome to out conference booking application!")
  fmt.Println("Get your tickets here to attend")


  var conferenceName = "Go Conference"
  fmt.Println("Conference Name: ", conferenceName, "booking-application")

  const conferenceTickets = 50
  fmt.Println("conferenceTickets",conferenceTickets)

  var remainingTickets = 100
  fmt.Println("remainingTickets",remainingTickets)
}