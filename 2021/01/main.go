package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
        log.Fatal(err)
    }

	contentAsString := string(content)
	contentAsStringArray := strings.Split(contentAsString, "\n")
	
	part1(contentAsStringArray)
	part2(contentAsStringArray)
}

func part1(contentAsStringArray []string) {
	var nbIncreaseMeasure = 0
	for i, v := range contentAsStringArray {
		if (i == 0) {
			fmt.Println(v, "(N/A - no previous measurement)")
			continue
		}

		prevValue, _ := strconv.Atoi(contentAsStringArray[i - 1]);
		currentValue, _ := strconv.Atoi(v)

		if (prevValue < currentValue) {
			fmt.Println(fmt.Sprintf("%v %v Increase \n", contentAsStringArray[i - 1], v))
			nbIncreaseMeasure++
		} else {
			fmt.Println(fmt.Sprintf("%v %v Decrease \n", contentAsStringArray[i - 1], v))
		}
	}
	fmt.Println(nbIncreaseMeasure)
}

func part2(contentAsStringArray []string) {
	var result = 0
	var prevSum = 0
	var currentSum = 0

	defer func() {
        if x := recover(); x != nil {
            fmt.Printf("Panic: %+v\n", x)
			fmt.Println(result)
        }
    }()

	for i, v := range contentAsStringArray {
		if (i == 0) {
			fmt.Println(v, "(N/A - no previous sum)")
			continue
		}

		currentValueIndex, _ := strconv.Atoi(v)
		prevValueIndex, _ := strconv.Atoi(contentAsStringArray[i-1])
		nextValueIndex, err := strconv.Atoi(contentAsStringArray[i+1])

		if err != nil {
			return 
		}

		nextNextValueIndex, err := strconv.Atoi(contentAsStringArray[i+2])

		if err != nil {
			return
		}

		prevSum = currentValueIndex + prevValueIndex + nextValueIndex
		currentSum = currentValueIndex + nextValueIndex + nextNextValueIndex

		if (prevSum < currentSum) {
			fmt.Println("Increase")
			result++
		}

		prevSum = 0
		currentSum = 0
	}
	
}