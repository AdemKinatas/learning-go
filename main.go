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

	whatwasSaid := saySomeThing()
	fmt.Println("The function returned", whatwasSaid)
}

func saySomeThing() string {
	return "something"
}
