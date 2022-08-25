package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var numOfPeople int
	var numOfPiecesEachPerson int
	var totalNumOfPieces int
	var requiredNumOfPizzas int
	var totalAmount int
	var lowestCost int
	var lowestCostOption string

	fmt.Println("Pizza Calculator")
	fmt.Println("How many number of people in the group?")
	fmt.Scanln(&numOfPeople)
	if numOfPeople <= 0 || numOfPeople > 999999999 {
		fmt.Println("Wrong Input.")
		os.Exit(0)
	}

	fmt.Println("How many pieces each person will eat? (Enter a value between 1 - 10)")
	fmt.Scanln(&numOfPiecesEachPerson)
	if numOfPiecesEachPerson <= 0 || numOfPiecesEachPerson > 10 {
		fmt.Println("Wrong Input. Please enter a value between 1-10")
		os.Exit(0)
	}

	totalNumOfPieces = numOfPeople * numOfPiecesEachPerson
	fmt.Println("Total number of pieces :", totalNumOfPieces)

	/*  5 inch Pizza */
	requiredNumOfPizzas = int(math.Ceil(float64(float64(totalNumOfPieces) / 4)))
	totalAmount = requiredNumOfPizzas * 3
	lowestCost = totalAmount
	lowestCostOption = "5 inch Pizzas X " + strconv.Itoa(requiredNumOfPizzas)
	fmt.Println("5 inch pizzas : $3 X ", requiredNumOfPizzas, " = $", totalAmount)

	/*  7 inch Pizza */
	requiredNumOfPizzas = int(math.Ceil(float64(float64(totalNumOfPieces) / 6)))
	totalAmount = requiredNumOfPizzas * 5
	if totalAmount != 0 && totalAmount < lowestCost {
		lowestCost = totalAmount
		lowestCostOption = "7 inch Pizzas X " + strconv.Itoa(requiredNumOfPizzas)
	}
	fmt.Println("7 inch pizzas : $5 X ", requiredNumOfPizzas, " = $", totalAmount)

	/*  10 inch Pizza */
	requiredNumOfPizzas = int(math.Ceil(float64(float64(totalNumOfPieces) / 10)))
	totalAmount = requiredNumOfPizzas * 8
	if totalAmount != 0 && totalAmount < lowestCost {
		lowestCost = totalAmount
		lowestCostOption = "10 inch Pizzas X " + strconv.Itoa(requiredNumOfPizzas)
	}
	fmt.Println("10 inch pizzas : $8 X ", requiredNumOfPizzas, " = $", totalAmount)

	/*  15 inch Pizza */
	requiredNumOfPizzas = int(math.Ceil(float64(float64(totalNumOfPieces) / 15)))
	totalAmount = requiredNumOfPizzas * 12
	if totalAmount != 0 && totalAmount < lowestCost {
		lowestCost = totalAmount
		lowestCostOption = "15 inch Pizzas X " + strconv.Itoa(requiredNumOfPizzas)
	}
	fmt.Println("15 inch pizzas : $12 X ", requiredNumOfPizzas, " = $", totalAmount)

	/*  24 inch Pizza */
	requiredNumOfPizzas = int(math.Ceil(float64(float64(totalNumOfPieces) / 24)))
	totalAmount = requiredNumOfPizzas * 18
	if totalAmount != 0 && totalAmount < lowestCost {
		lowestCost = totalAmount
		lowestCostOption = "24 inch Pizzas X " + strconv.Itoa(requiredNumOfPizzas)
	}
	fmt.Println("24 inch pizzas : $8 X ", requiredNumOfPizzas, " = $", totalAmount)

	fmt.Println("Lowest cost option : ", lowestCostOption)
	fmt.Println("Lowest cost option price: $", lowestCost)

}
