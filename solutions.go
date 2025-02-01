package main

import "fmt"

func main(){
	
	var array = []int{3,2,4}
	fmt.Print(twoSum(array,6))
}

func twoSum(array1 []int, target int ) []int{
	// var hashmap = make(map[int]int)
	// for i := 0; i < len(array1);i++{
	// 	var difference int = target - array1[i];
	// 	// value, ok:= hashmap[difference]
	// 	if _, exists:=hashmap[difference];exists{
	// 		nums := []int{hashmap[difference],i}
	// 		return nums
	// 	}
	// 	hashmap[array1[i]] = i

	// }

	// return []int{}

	var hashmap = make(map[int]int)
	for i := 0; i < len(array1); i++{
		difference := target - array1[i]
		if _, exists := hashmap[difference];exists{
			return []int{hashmap[difference],i}
		}
		hashmap[array1[i]] = i
	}

	return []int{}
}

func duplicate([]) int{

}