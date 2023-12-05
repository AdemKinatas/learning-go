package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	Lastname  string
}

func main() {
	stringMap := make(map[string]string)

	stringMap["dog"] = "samson"
	stringMap["other-dog"] = "cassie"

	// log.Println(stringMap["dog"], stringMap["other-dog"])

	intMap := make(map[string]int)

	intMap["first"] = 1
	intMap["Second"] = 2

	// log.Println(intMap["first"], intMap["Second"])

	userMap := make(map[string]User)

	me := User{
		FirstName: "Adem",
		Lastname:  "Kınataş",
	}

	userMap["me"] = me

	// log.Println(userMap["me"].FirstName)

	var stringSlice []string

	stringSlice = append(stringSlice, "Adem")
	stringSlice = append(stringSlice, "John")

	// log.Println(stringSlice)

	var intSlice []int

	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 3)

	sort.Ints(intSlice)

	// log.Println(intSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// log.Println(numbers)
	log.Println(numbers[6:9])

	names := []string{"one", "two", "cat", "fish"}
	log.Println(names)
}
