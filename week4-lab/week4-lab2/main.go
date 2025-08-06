package main

import (
	"fmt"
)

// var email string = "onwan_k@su.ac.th"

func main() {
	//var name string = "kessara"
	var age int = 21

	email := "onwan_k@su.ac.th"
	gpa := 4.00

	firstname, lastname := "kessara", "onwan"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)

}
