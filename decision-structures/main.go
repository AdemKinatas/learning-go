package main

import (
	"log"
)

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		// log.Println("isTrue is", isTrue)
	} else {
		// log.Println("isTrue is", isTrue)
	}

	isTrue = false
	myNum := 100

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 || isTrue {
		log.Println("1")
	} else if myNum == 101 || isTrue {
		log.Println("2")
	} else if myNum > 1000 && isTrue == false {
		log.Println("3")
	}

	animal := "horse"

	switch animal {
	case "cat":
		log.Println("animal is set to cat")
	case "dog":
		log.Println("animal is set to dog")
	case "fish":
		log.Println("animal is set to fish")
	default:
		log.Println("animal is something else")
	}
}
