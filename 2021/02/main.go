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
	var posX = 0
	var depth = 0

	for _, v := range contentAsStringArray {
		keyValue := strings.Split(v, " ")

		action := keyValue[0]
		value, _ := strconv.Atoi(keyValue[1])

		if (action == "forward") {
			posX += value
			continue
		} 

		if (action == "up") {
			depth -= value
			continue
		}

		if (action == "down") {
			depth += value
			continue
		}
	}

	fmt.Println("posX", posX)
	fmt.Println("depth", depth)

	finalResult := posX * depth

	fmt.Println(finalResult)
}

func part2 (contentAsStringArray []string) {
	var posX = 0
	var depth = 0
	var aim = 0

	for _, v := range contentAsStringArray {
		keyValue := strings.Split(v, " ")

		action := keyValue[0]
		value, _ := strconv.Atoi(keyValue[1])

		if (action == "forward") {
			posX += value
			depth += aim * value
			continue
		} 

		if (action == "up") {
			aim -= value
			continue
		}

		if (action == "down") {
			aim += value
			continue
		}
	}

	fmt.Println("posX", posX)
	fmt.Println("depth", depth)

	finalResult := posX * depth

	fmt.Println(finalResult)

}