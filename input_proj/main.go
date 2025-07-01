package main

import (
	"bufio"
	"fmt"
	"os"
)

// 2 task
/*func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	inputStr := scanner.Text()
	fmt.Println(inputStr, "- Лучшая книга!")
	// var num int
	// _, err := fmt.Scan(&num)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
}*/

// 3 task
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	firstStr := scanner.Text()
	_ = scanner.Scan()
	secondStr := scanner.Text()
	_ = scanner.Scan()
	thirdStr := scanner.Text()
	fmt.Println(firstStr)
	fmt.Println(secondStr)
	fmt.Println(thirdStr)
}
