package main

import "fmt"

func getFrequency(num int, arr []int) int {
	var freq int = 0
	for _, value := range arr {
		if num == value {
			freq = freq + 1
		}
	}
	return freq
}

func part2() int {
	fmt.Println(1 == 1)
	firstArr, secondArr := getLists()
	var totalDistance int = 0
	for _, value := range firstArr {
		freq := getFrequency(value, secondArr)
		distance := value * freq
		totalDistance = totalDistance + distance
	}
	return totalDistance
}
