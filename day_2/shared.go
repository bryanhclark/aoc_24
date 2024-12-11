package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLines() [][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var allLines [][]string
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Print(line)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		numArr := strings.Split(strings.TrimSuffix(line, "\n"), " ")
		allLines = append(allLines, numArr)
	}
	return allLines
}
