package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLists() ([]int, []int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var firstArr []int
	var secondArr []int
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		splitResult := strings.Split(strings.TrimSuffix(line, "\n"), "   ")
		first, firstErr := strconv.Atoi(splitResult[0])
		second, secondErr := strconv.Atoi(splitResult[1])
		if firstErr != nil {
			log.Fatal(secondErr)
		}
		if secondErr != nil {
			log.Fatal(secondErr)
		}
		firstArr = append(firstArr, first)
		secondArr = append(secondArr, second)
	}
	return firstArr, secondArr
}
