package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var target = 2020
var nbOfDeterminants = 3

func main () {
	
    content, err := ioutil.ReadFile("input.txt")
	contentAsString := string(content)
	contentAsStringArray := strings.Split(contentAsString, "\n")

	findTwoNumbers(contentAsStringArray)
	findThreeNumbers(contentAsStringArray)
	
    if err != nil {
        log.Fatal(err)
    }

}

func findTwoNumbers(contentAsStringArray []string) {
	for _, i := range contentAsStringArray {
		firstDeterminant, _ := strconv.Atoi(i)
		for _, j := range contentAsStringArray {
			secondDeterminant, _ := strconv.Atoi(j)
			if (firstDeterminant + secondDeterminant == target) {
				fmt.Println("Determinants are", firstDeterminant, secondDeterminant)
				fmt.Println("Result is", firstDeterminant*secondDeterminant)
				return
			}
		}
	}
}

func findThreeNumbers(contentAsStringArray []string) {
	for _, i := range contentAsStringArray {
		firstDeterminant, _ := strconv.Atoi(i)
		for _, j := range contentAsStringArray {
			secondDeterminant, _ := strconv.Atoi(j)

			for _, x := range contentAsStringArray {
				thirdDeterminant, _ := strconv.Atoi(x)
				if (firstDeterminant + secondDeterminant + thirdDeterminant == target) {
					fmt.Println("Determinants are", firstDeterminant, secondDeterminant, thirdDeterminant)
					fmt.Println("Result is", firstDeterminant*secondDeterminant*thirdDeterminant)
					return
				}
			}

		}
	}
}