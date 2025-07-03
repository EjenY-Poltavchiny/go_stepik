package main

import "fmt"

func main() {
	var firstInput int
	var secondInput int
	var thirdInput int
	_, err := fmt.Scan(&firstInput, &secondInput, &thirdInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(firstInput * secondInput * thirdInput)
}
