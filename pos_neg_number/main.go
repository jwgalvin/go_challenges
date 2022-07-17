package main

import "fmt"

func main(){
	inte := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(inte))
}
func CountPositivesSumNegatives(numbers []int) []int {
	result := make([]int, 2, 2)
	for _, pn := range numbers {
		if pn > 0 {
			result[0]++
		}else{
			result[1] = result[1] + pn
		}	
	}
	return result
}