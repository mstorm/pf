package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	. "strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s file line ...\n", path.Base(os.Args[0]))
		os.Exit(-1)
	}

	targetFilename := os.Args[1]
	targetFile, err := os.Open(targetFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer targetFile.Close()

	var lines []int
	for _, lineStr := range os.Args[2:] {
		line, err := Atoi(lineStr)
		if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, line)
	}
	sort.Ints(lines)

	var currentLine = 0
	scanner := bufio.NewScanner(targetFile)
	for scanner.Scan() {
		if len(lines) < 1 {
			break
		}

		currentLine += 1

		if currentLine == lines[0] {
			fmt.Println(scanner.Text())
			lines = lines[1:]
		}
	}
}
