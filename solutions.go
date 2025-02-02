package main

import (
	"fmt"
	"math"
)

func main() {

	var array = []int{3, 2, 4}
	fmt.Println(twoSum(array, 6))
	fmt.Println(reverse_integer(23))
}

func twoSum(array1 []int, target int) []int {

	var hashmap = make(map[int]int)
	for i := 0; i < len(array1); i++ {
		difference := target - array1[i]
		if _, exists := hashmap[difference]; exists {
			return []int{hashmap[difference], i}
		}
		hashmap[array1[i]] = i
	}

	return []int{}
}

func reverse_integer(number int) bool {
	var reverse_int int
	//store the original number to be compared
	var original_number int = number
	for number != 0 {
		//modulo of the number is
		var digit int = number % 10

		//Divide number with 10
		number = number / 10

		//If overflow happened
		if reverse_int > math.MaxInt32 || reverse_int < math.MinInt32 {
			fmt.Print("Overflow")
			return false
		}
		reverse_int = reverse_int*10 + digit
	}
	if reverse_int == original_number {
		return true
	} else {
		return false
	}
}
