package main

import "log"

func main() {
	var myString string

	myString = "Green"

	log.Println("mtString is set to ", myString)
	changeUsingPointer((&myString))
	log.Println("after function call mtString is set to ", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to ", s)
	newValue := "Red"
	*s = newValue
}
