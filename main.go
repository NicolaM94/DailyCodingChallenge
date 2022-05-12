package main

import (
	"fmt"
	"dcc/utilities"
)

// Given an array of integers and a target integer k, return true if the array
// contains two numbers which sum up to k.
// 11/05/2022
func findSumInArray (array []int, kvalue int) bool {

	for _,number := range(array) {
		if  utilities.Contains(array,kvalue-number) {
			return true
		}
	}
	return false
}

// 12/05/2022
// Given an array of integers return a new array such that every element at index i
// of the new array is the product of all the numbers in the original array except the one at i.
// Pro: What if you can't divide?

func productExceptOne (arrayOfInts []int) []int {

	var output []int
	if utilities.Contains(arrayOfInts,0) {
		for i := 0; i < len(arrayOfInts)-1; i++ {
			output = append(output, 0)
		}
		return output
	}
	for _,value := range(arrayOfInts) {
		output = append(output, utilities.Products(arrayOfInts)/value)
	}
	return output
}



func main () {
	fmt.Println(findSumInArray([]int{1,2,3,4,5,6,7},7))
	fmt.Println(productExceptOne([]int{0,1,2,3,4,5,6}))
}