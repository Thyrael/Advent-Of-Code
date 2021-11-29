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
	contentAsString := string(content)
	contentAsStringArray := strings.Split(contentAsString, "\n")
	
	partOne(contentAsStringArray)
	partTwo(contentAsStringArray)


	if err != nil {
        log.Fatal(err)
    }
}

func partOne(contentAsStringArray []string) {
	var nbOfValidPwd = 0

	for _, v := range contentAsStringArray {
		minLimit, maxLimit, char, input := extractDataFromLine(v)

		countChar := strings.Count(input, char)

		if (minLimit <= countChar && countChar <= maxLimit) {
			fmt.Println("OK", char, input)
			nbOfValidPwd++
		}
	}

	fmt.Println(nbOfValidPwd)
}

func extractDataFromLine(input string) (int,int, string, string) {
	t := strings.Split(input, " ")
	minLimit, maxLimit := calcLimit(t[0])
	char := strings.Split(t[1], ":")[0]

	return minLimit,maxLimit, char, t[2]
}

func calcLimit(limitAsString string) (int,int) {
	limits := strings.Split(limitAsString, "-")

	minLimit, _ := strconv.Atoi(limits[0])
	maxLimit, _ := strconv.Atoi(limits[1])

	return minLimit, maxLimit
} 

func partTwo(contentAsStringArray []string) {
	var nbOfValidPwd = 0

	for _, v := range contentAsStringArray {
		minLimit, maxLimit, char, input := extractDataFromLine(v)
		
		var minLimitOK = false
		var maxLimitOK = false
	
		for i, rune := range input {
			i++

			if (i == minLimit && string(rune) == char) {
				fmt.Println("minLimit char is", string(rune))
				minLimitOK = true
				continue
			}

			if (i == maxLimit && string(rune) == char) {				
				fmt.Println("maxLimit char is", string(rune))
				maxLimitOK = true
				continue
			}
		}

		if (minLimitOK && !maxLimitOK || !minLimitOK && maxLimitOK) {
			nbOfValidPwd++
		}
	}

	fmt.Println(nbOfValidPwd)
}