package main

import "fmt"

func main() {
	var list = make(map[string]struct{})
	p := &list
	substrFunc(p, "abcde")
	fmt.Println("Number of substrings = ", len(list))

	for i, _ := range *p {
		fmt.Println(i)
	}
}

func substrFunc(subStrList *map[string]struct{}, inputStr string) {
	(*subStrList)[inputStr] = struct{}{}
	var length = len(inputStr)

	for i := 1; i < length; i++ {
		subStr := inputStr[i:length]
		substrFunc(subStrList, subStr)
		subStr = inputStr[0 : length-i]
		substrFunc(subStrList, subStr)
	}
}
