package main

import "fmt"

func main(){
	inte := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(inte))
}
func CountPositivesSumNegatives(numbers []int) []int {
	var pos, neg int
	result := make([]int, 2, 2)
	for _, pn := range numbers {
		if pn > 0 {
			pos++
		}else{
			neg = neg + pn
		}
		result[0] = pos
		result[1] = neg
	}
	return result
}