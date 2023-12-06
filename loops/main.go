package main

import (
	"log"
)

func main() {
	// for loop
	for i := 0; i <= 10; i++ {
		// log.Println(i)
	}

	animals := []string{"dog", "cat", "horse", "fish", "cow"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	log.Println("****************************************************************************************************")

	mapAnimals := make(map[string]string)
	mapAnimals["dog"] = "Fido"
	mapAnimals["cat"] = "Fluffy"

	for animalType, animal := range mapAnimals {
		log.Println(animalType, animal)
	}

	log.Println("****************************************************************************************************")

	var firstLine = "Once upon a time at the America"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	log.Println("****************************************************************************************************")

	type User struct {
		Firstname string
		Lastname  string
		Email     string
		Age       int
	}

	var users []User

	users = append(users, User{Firstname: "Adem", Lastname: "Kınataş", Email: "adem@gmail.com", Age: 31})
	users = append(users, User{Firstname: "Onur", Lastname: "Kınataş", Email: "onur@gmail.com", Age: 24})
	users = append(users, User{Firstname: "Enes", Lastname: "Kınataş", Email: "enes@gmail.com", Age: 31})

	for _, user := range users {
		log.Println(user.Firstname, user.Lastname, user.Email, user.Age)
	}
}
