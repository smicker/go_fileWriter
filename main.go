package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var file string = "./output.txt"

func main() {
	fmt.Println("Starting program")

	timeInterval := time.Second * 1

	i := 0
	for {
		fmt.Printf("Writing %d to %s\n", i, file)
		if err := os.WriteFile(file, []byte(strconv.Itoa(i)), 0644); err != nil {
			panic(err)
		}
		time.Sleep(timeInterval)
		i += 1
	}
}
