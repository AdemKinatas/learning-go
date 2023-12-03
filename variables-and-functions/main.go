package main

import "fmt"

func main() {
	fmt.Println("Hello world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye cruel world!"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid := saySomeThing()
	fmt.Println("The function returned", whatWasSaid)

	whatHeWasSaid, whatSheWasSaid := multiSaySomeThings()
	fmt.Println("The second function returned", whatHeWasSaid, whatSheWasSaid)
}

func saySomeThing() string {
	return "something"
}

func multiSaySomeThings() (string, string) {
	return "something", "else"
}
