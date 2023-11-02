package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func wordCount(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to readfile: %v", err)
	}

	defer f.Close()

	fScanner := bufio.NewScanner(f)

	fScanner.Split(bufio.ScanWords)

	count := 0

	for fScanner.Scan() {
		count++
	}

	if err := fScanner.Err(); err != nil {
		panic(err)
	}

	return count
}

func main() {
	filepath := os.Args[1]
	if len(os.Args) < 2 {
		fmt.Println("Program accepts 2")
	}
	fmt.Println(wordCount(filepath))
}
