package main

import (
	"fmt"
)

func main() {
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	fmt.Printf("marks=%d, grade=%s\n", marks, grade)

	switch {
	case grade == "A":
		fmt.Println("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Println("Well Done!\n")
	case grade == "D":
		fmt.Println("You passed\n")
	case grade == "F":
		fmt.Println("Better try again\n")
	default:
		fmt.Println("Invalid grade\n")
	}
}
