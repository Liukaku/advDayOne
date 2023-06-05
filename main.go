package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	fileOpen, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("file err", err.Error())
	}

	fileRead, err := ioutil.ReadAll(fileOpen)

	calories := string(fileRead)

	calSlice := strings.Split(calories, "\n")

	var calCount int64
	highestCal := []int64{ 0,0,0 }

	fmt.Print(highestCal)
	for _, elf := range calSlice {
		if len(strings.TrimSpace(elf)) == 0 {
			appended := false
			for n, val := range highestCal {

				if val < calCount && !appended{
					newHighest := InsertIntoSliceAtIndex(highestCal, calCount, n)
					highestCal = append(newHighest[:3])
					appended = true
				}
			}
			appended = false
			calCount = 0
		} else {
			val, err := strconv.ParseInt(elf, 0, 0)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("error parsing", elf)
			}

			calCount += val
		}
	}

	var totalCount int64 = 0

	for _, val := range highestCal {
		totalCount += val
	}

	fmt.Println("highest: ", highestCal)
	fmt.Println("count: ", totalCount)
}

func InsertIntoSliceAtIndex[T any](destination []T, element T, index int) []T {
	if len(destination) == index {
		return append(destination, element)
	}

	destination = append(destination[:index+1], destination[index:]...)
	destination[index] = element

	return destination
}