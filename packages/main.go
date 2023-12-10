package main

import (
	"log"

	"github.com/tsawler/myniceprogram/helpers"
)

func main() {
	log.Println("Hello!")

	var mayVar helpers.SomeType
	mayVar.TypeName = "some name"
	log.Println(mayVar.TypeName)
}
