package main

import (
	"fmt"
	"math"
	"strings"
)




func main() {

	var array = []int{3, 2, 4}
	fmt.Println(twoSum(array, 6))
	fmt.Println(reverse_integer_check(120))
	fmt.Println(roman_to_integer("LVIII"))
	fmt.Println(valid_word("234Adas"))

	//Instantiate the Stack struct
	fmt.Println(valid_parentheses("([])"))

	
}

func twoSum(array1 []int, target int) []int {

	//Create a map
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

func reverse_integer_check(number int) bool {
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
		fmt.Println(reverse_int)
		return false
	}
}

func roman_to_integer(romanNumber string) int {
	//Map for the roman numerals
	var romanMap = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	//Actual number
	var total int = 0
	for i := 0; i < len(romanNumber); i++ {
		//if the next string doesnt pass the length of romanNumber and value of the romanNumber from the romanMap key is currently less than the next value
		if i+1 < len(romanNumber) && romanMap[string(romanNumber[i])] < romanMap[string(romanNumber[i+1])] {
			total -= romanMap[string(romanNumber[i])]
		} else {
			total += romanMap[string(romanNumber[i])]
		}
	}
	return total
}

func valid_word(word string) (bool,bool){
	var vowel string = "aiueo"
	vowel +=  strings.ToUpper(vowel)
	var consonant string = "bcdfghjklmnpqrstvwxyz"
	consonant += strings.ToUpper(consonant)
	var numbers string = "012346789"
	var allowed_string string = vowel + consonant + numbers 
	//Status check
	var status_vowel bool = false
	var status_consonant bool = false

	//check if strings is no longer than 3 character
	if len(word) < 3{
		return false,false
	}

	for i:=0; i < len(word); i++{
		if strings.Contains(vowel,string(word[i])){
			status_vowel = true
		}
		if strings.Contains(consonant,string(word[i])){
			status_consonant = true
		}
		if !strings.Contains(allowed_string,string(word[i])){
			return false,false
		}
	}
	return status_vowel,status_consonant
}




func valid_parentheses(s string) bool{
	var close_to_open = map[string]string{"}":"{","]":"[",")":"("}
	stack := Stack{}
	//iterate through the string
	for i:=0; i < len(s);i++{
		//jika substring adalah key didalam close_to_open map
		if _, exist := close_to_open[string(s[i])]; exist{
			if stack.Size() > 0 && stack.Top() == close_to_open[string(s[i])]{
				stack.Pop()
			}else{
				return false
			}
		}else{
			stack.Push(string(s[i]))
		}
	}
	return true

}