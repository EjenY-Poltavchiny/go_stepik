package main

import (
	"fmt"
	"math"
)

// math 1
// func main() {
//     var aVal float64
//     var bVal float64
//     fmt.Scan(&aVal, &bVal)
//     fmt.Println(aVal + bVal + math.Sqrt(math.Pow(aVal, 2) + math.Pow(bVal, 2)))
// }

// math 3
// func main() {
// 	var x1 float64
// 	var x2 float64
// 	var y1 float64
// 	var y2 float64
// 	fmt.Scan(&x1, &y1, &x2, &y2)
// 	fmt.Println(math.Abs(x1-x2) + math.Abs(y1-y2))
// }

// test
// func main() {
// 	var aVal int
// 	fmt.Scan(&aVal)
// 	fmt.Println((aVal%1000)/100 + (aVal%100)/10 + aVal%10)
// }

func main() {
	var aVal float64
	var bVal float64
	fmt.Scan(&aVal, &bVal)
	fmt.Println(math.Sqrt(math.Pow(aVal, 2) + math.Pow(bVal, 2)))
}
