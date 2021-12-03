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
	
	//part1(contentAsStringArray)
	part2(contentAsStringArray)
}


func part1(contentAsStringArray []string) {
	
	lineLen := len(contentAsStringArray[0])

	var gammaRateAsBinary = ""
	var epsilonRateAsBinary = ""

	for i := 0; i < lineLen; i++ {
		oneBit, zeroBit := iterateEachChar(contentAsStringArray, i)
		if (oneBit > zeroBit) {
			gammaRateAsBinary += "1"
			epsilonRateAsBinary += "0"
		} else {
			gammaRateAsBinary += "0"
			epsilonRateAsBinary += "1"
		}
	}


	epsilonRate, _ := strconv.ParseInt(epsilonRateAsBinary, 2, 64) 

	fmt.Println("as binary", epsilonRateAsBinary)
	fmt.Println("as decimal ", epsilonRate)

	gammaRate, _ := strconv.ParseInt(gammaRateAsBinary, 2, 64) 

	fmt.Println("as binary", gammaRateAsBinary)
	fmt.Println("as decimal ", gammaRate)

	consumption := epsilonRate * gammaRate

	fmt.Println(consumption)

}

func iterateEachChar(contentAsStringArray []string, position int) (int,int) {
	var oneBit = 0
	var zeroBit = 0

	for _, v := range contentAsStringArray {
		if(string(v[position]) == "1") {
			oneBit++
		} else {
			zeroBit++
		}	
	}

	return oneBit, zeroBit
}


func part2 (contentAsStringArray []string) {
	
	oxygenRateAsBinary := findOxygenRate(contentAsStringArray)
	oxygenRate, _ :=  strconv.ParseInt(oxygenRateAsBinary, 2, 64)
	fmt.Println("Oxygen rate", oxygenRate)

	co2Rate, _ := strconv.ParseInt(findCo2Rate(contentAsStringArray), 2, 64) 
	fmt.Println("Co2 rate", co2Rate)
	
	finalResult := oxygenRate * co2Rate
	fmt.Println(finalResult)
}

func findCo2Rate(co2RatingValues []string) string { 
	lineLen := len(co2RatingValues[0])

	for i := 0; i < lineLen; i++ {
		// fmt.Println("co2RatingValues", co2RatingValues)
		oneBit, zeroBit := iterateEachChar(co2RatingValues, i)
		
		if (oneBit >= zeroBit) {
			delimiter := "0"
		
			listOfIndexToRemove := filteredIndexArray(co2RatingValues, delimiter, i)
		
			for _ , v := range listOfIndexToRemove {
				co2RatingValues = append(co2RatingValues[:v], co2RatingValues[v+1:]...)
			}
		} else {
			delimiter := "1"
			listOfIndexToRemove := filteredIndexArray(co2RatingValues, delimiter, i)
			
			for _ , v := range listOfIndexToRemove {
				co2RatingValues = append(co2RatingValues[:v], co2RatingValues[v+1:]...)
			}
		}

		if(len(co2RatingValues) == 1) {break}

	}

	fmt.Println("Co2 rate binary ", co2RatingValues)
	return co2RatingValues[0]
}

func findOxygenRate(oxygenRatingValues []string) string {
	lineLen := len(oxygenRatingValues[0])

	for i := 0; i < lineLen; i++ {
		// fmt.Println(oxygenRatingValues)
		oneBit, zeroBit := iterateEachChar(oxygenRatingValues, i)
		
		if (oneBit >= zeroBit) {
			delimiter := "1"
		
			listOfIndexToRemove := filteredIndexArray(oxygenRatingValues, delimiter, i)
		
			for _ , v := range listOfIndexToRemove {
				oxygenRatingValues = append(oxygenRatingValues[:v], oxygenRatingValues[v+1:]...)
			}
		} else {
			delimiter := "0"
			listOfIndexToRemove := filteredIndexArray(oxygenRatingValues, delimiter, i)
			
			for _ , v := range listOfIndexToRemove {
				oxygenRatingValues = append(oxygenRatingValues[:v], oxygenRatingValues[v+1:]...)
			}
		}
		if(len(oxygenRatingValues) == 1) {break}
	}

	fmt.Println("Oxygen rate binary ", oxygenRatingValues)
	return oxygenRatingValues[0]
}

func filteredIndexArray(contentAsStringArray []string, delimiter string, position int) []int {
	var tmp []int
	for i, v := range contentAsStringArray {
		if(string(v[position]) != delimiter) {
			tmp = append(tmp, i)
		}
	}

	return reverseFilteredIndexArray(tmp)
}

func reverseFilteredIndexArray(filterIndexArray []int) []int {
	reversed := []int{}
        
    for i := range filterIndexArray {
        n := filterIndexArray[len(filterIndexArray)-1-i]
        reversed = append(reversed, n)
    }

	return reversed
}