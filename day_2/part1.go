package main

import (
	"fmt"
	"math"
	"strconv"
)

// TODO:
// - add increase/descreasing checks
// - check for same number
func isSafe(arr []string) bool {
	var isSafe bool = true
	isDescending := arr[0] > arr[1]
	for i := 1; i < len(arr); i++ {
		curr, _ := strconv.Atoi(arr[i])
		prev, _ := strconv.Atoi(arr[i-1])
		difference := math.Abs(float64(curr) - float64(prev))
		if isDescending {
			// if we're descending, but the curr is greater
			if curr > prev {
				isSafe = false
			}
		} else {
			// is we're ascending, and the curr is smaller than prev
			if prev > curr {
				isSafe = false
			}
		}

		isSameValue := prev == curr
		if difference > 3 || isSameValue {
			isSafe = false
		}
	}
	return isSafe
}

func part1() {
	lines := getLines()
	var safeCount int = 0
	for _, val := range lines {
		isSafeRes := isSafe(val)
		if isSafeRes {
			safeCount++
		}
	}
	fmt.Println("safeCount: ", safeCount)
}
