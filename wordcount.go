package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("incorrect input. usage: ./wordcount 'text'")
		os.Exit(1)
	}
	words := os.Args[1]
	split := bytes.Split([]byte(words), []byte(" "))
	count := 0
	for _, i := range split {
		str := string(i)
		if str != "" {
			count++
		}
	}
	fmt.Println(count)
}
