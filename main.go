package main

import (
	"dcc/utilities"
	"fmt"
	
)

// 11/05/2022
// Given an array of integers and a target integer k, return true if the array
// contains two numbers which sum up to k.

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


// 14/05/2022
// Given an array of integers, find the first missing positive integer in linear time and constant space.
// In other words, find the lowest positive integer that does not exist in the array.
// The array can contain duplicates and negative numbers as well.

func missingPositive (integers []int) int {
	min := utilities.FindMinimun(integers)
	if min == 0 {
		return 1
	} else if min == 1 {
		return 0
	} else {
		return min-1
	}
}


func main () {
	fmt.Println(findSumInArray([]int{1,2,3,4,5,6,7},7))
	fmt.Println(productExceptOne([]int{0,1,2,3,4,5,6}))
	fmt.Println(missingPositive([]int{40,3,5,-1,4,4,9}))
}