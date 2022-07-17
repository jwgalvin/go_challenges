package main

import "fmt"

func main(){
	inte := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(inte))
}
func CountPositivesSumNegatives(numbers []int) []int {
	var pos int
	var neg int
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
	// length := len(numbers)
	// out := []int{}
	
	// if length == 0{
	// 	return out
	// }else {
	// 	out := []int{}
	// 	for i, pn := range numbers{
	// 		if pn >= 0 {
	// 			pos++
	// 			out[0] = pos
	// 	}else{
	// 			neg = neg + numbers[i]
	// 			out[1] = neg
	// 		}
			
	// 	}
	// }
	// return fmt.Println(out)

}