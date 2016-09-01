package main

import(
	"fmt"
	"math/rand"
)

func main() {
	var text1, text2 string
	text1 = "test"
	text2 = "string"
	text1, text2 = returnTwoStrings(text1, text2)
	// rand.Seed(time.Now().Millisecond())
	fmt.Println("Random number and output ", rand.Int())
	fmt.Println("combining a string: ",text1, text2)
	fmt.Println(threes(100));
}

func threes( original int) int{
	fmt.Println("Input: ", original)
	var currentValue int = original

	for currentValue!= 1 {
		
		if currentValue%3 == 0 {
			fmt.Println(currentValue, " 0")
			currentValue = currentValue/3
		} else if currentValue%3 == 1 {
			fmt.Println(currentValue, " -1")
			currentValue = (currentValue-1)/3
		} else if currentValue%3 == 2 {
			fmt.Println(currentValue, " +1")
			currentValue = (currentValue+1)/3
		}
		
	}

	return currentValue;
}

//go can return multiple variables
func returnTwoStrings(a, b string) (string, string) {
	return (a+b), (b+a)
}



