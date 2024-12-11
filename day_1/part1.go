package main

import (
	"fmt"
	"math"
	"sort"
)

func calculateDistance(arr1 []int, arr2 []int) float64 {
	var totalDistance float64 = 0
	for index, _ := range arr1 {
		first := arr1[index]
		second := arr2[index]
		distance := math.Abs(float64(first) - float64(second))
		totalDistance = totalDistance + distance
	}
	return totalDistance
}

func part1() {
	firstArr, secondArr := getLists()
	sort.Ints(firstArr)
	sort.Ints(secondArr)
	distance := calculateDistance(firstArr, secondArr)
	fmt.Println("totalDistance: ", int(distance))
}
