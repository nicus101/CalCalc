package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		menu     string
		distance float64
		actsize  float64
		dessize  float64
	)

	fmt.Println("Welcome to CalCalc")
	fmt.Println("Select your firmware \n M for Marlin, K for Klipper")

	_, err := fmt.Scanf("%s", &menu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("odczytano %s \n", menu)
	if menu == "m" {
		fmt.Println("Enter steps per mm value")
		_, err = fmt.Scanf("%f", &distance)
		fmt.Println("Enter desired length")
		_, err = fmt.Scanf("%f", &dessize)
		fmt.Println("Enter actual length")
		_, err = fmt.Scanf("%f", &actsize)

		fmt.Printf("distance: %f \ndesired length: %f \nactual length %f \n", distance, dessize, actsize)

	}
}
