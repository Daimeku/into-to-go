package main

import(
	"fmt"
	"strings"
	"math/rand"
    //"golang.org/x/tour/pic"
)

//types seem similar to classes
type Headphone struct{
	brand, model string
}

func main() {
	var text1, text2 string
	text1 = "test"
	text2 = "string"
	text1, text2 = returnTwoStrings(text1, text2)
	
	fmt.Println("Random number and output ", rand.Int())
	fmt.Println("combining a string: ",text1, text2)
	fmt.Println(threes(100));

	testMap()
	//from go tutorial
	//pic.Show(Pic)
}

//from go tour excercise-maps.go
//splits a string, counts the words and returns a map[string]int with the wordcount
func WordCount(s string) map[string]int {
	var splitString []string = strings.Fields(s) //split the string into a slice of substrings
	var countMap map[string]int = make(map[string]int)

	//loop through the slice of substrings and count the occurrences of each string
	for i:=0;i<len(splitString);i++ {

		if val, ok := countMap[splitString[i]]; ok == true { //check if the value is in the map- this funcion returns 2 values
			countMap[splitString[i]] = val+1 //if the value is in the map then increment it
		} else {
			countMap[splitString[i]] = 1 // if the value isnt in the map then initialize it to 1
		}
	}

	return countMap
}


//maps store key/value pairs
func testMap() {

	//create a map literal of type Headphone
	var headphones = map[string]Headphone{
		"jbl" : Headphone{ 
			"JBL", "Reflect Response",
			},
	}
	fmt.Println("map literal: ", headphones["jbl"])


	headphones = make(map[string]Headphone) //make a map of type Headphone
	headphones["hd449"] = Headphone{ "Sennheiser", "HD449"} //create a Headphone instance and assign it to the "hd449" key in the map
	
	fmt.Println("map output: ",headphones["hd449"])
}

//https://www.reddit.com/r/dailyprogrammer/comments/3r7wxz/20151102_challenge_239_easy_a_game_of_threes/
func threes( original int) int{
	fmt.Println("Input: ", original)
	var currentValue int = original

	//until the current value is 3
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

//from the go tutorial, created a slice of a slice
// picture is a slice []uint8 containing picRow []uint8
func Pic(dx, dy int) [][]uint8 {
	
	picture := make([][]uint8, dy)
	
	for i:=0;i<len(picture);i++ {
		picRow := make([]uint8, dx)
		
		for y:= 0;y<len(picRow)-1;y++ {
			picRow[y] = uint8((i*y)/2)
		}
		picture[i] = picRow
	}
	return picture
}



