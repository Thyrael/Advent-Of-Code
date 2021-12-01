package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	contentAsString := string(content)
	contentAsStringArray := strings.Split(contentAsString, "\n")


	nbOfThreeFirstSlope := calcThreeInSlope("slope1_1", 1, 1, contentAsStringArray)
	nbOfThreeSecondSlope := calcThreeInSlope("slope3_1", 3, 1, contentAsStringArray)
	nbOfThreeThirdlope := calcThreeInSlope("slope5_1", 5, 1, contentAsStringArray)
	nbOfThreeFourSlope := calcThreeInSlope("slope7_1", 7, 1, contentAsStringArray)
	nbOfThreeFiveSlope := calcThreeInSlope("slope1_2", 1, 2, contentAsStringArray)

	fmt.Println(nbOfThreeFiveSlope)

	total := nbOfThreeFirstSlope * nbOfThreeSecondSlope * nbOfThreeThirdlope * nbOfThreeFourSlope * nbOfThreeFiveSlope

	fmt.Println(total)

	if err != nil {
        log.Fatal(err)
    }
}

func calcThreeInSlope(slopeId string, slopeX int, slopeY int, contentAsStringArray []string) (result int) {

	var initialX = 0
	var initialY = 0

	var nbOfThree = 0

	defer func() {
        if x := recover(); x != nil {
            fmt.Printf("Panic: %+v\n", x)
			result = nbOfThree
        }
    }()

	f, _ := os.Create(fmt.Sprintf("output_%v.txt", slopeId))
    defer f.Close()


	rowLen := len(contentAsStringArray[0])

	
	for _, _ = range contentAsStringArray {
		value := contentAsStringArray[initialY]

		for i , rune := range value {
			if (i == initialX) {
				// fmt.Println(string(rune))
				if (string(rune) == "#") {
					nbOfThree++
					f.WriteString("X")
				} else {
					f.WriteString("0")
				}
			} else {
				f.WriteString(string(rune))
			}
		}
		f.WriteString("\n")

		if (initialX + slopeX >= rowLen) {
			report := rowLen - initialX
			initialX = slopeX - report
		} else {
			initialX += slopeX
		}
		initialY += slopeY
	}

	fmt.Println(nbOfThree)
	result = nbOfThree
	return result
}