package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var maxIter int
	var crashIter int
	var fileName string
	flag.IntVar(&maxIter, "runTime", 0, "Makes program exit ok after this amount of seconds. 0 (default) means infinite.")
	flag.IntVar(&crashIter, "crashTime", 0, "Makes program crash after this amount if seconds. 0 (default) means infinite.")
	flag.StringVar(&fileName, "fName", "./output.txt", "Filename for the output.")
	flag.Parse()

	fmt.Println("Starting program")

	timeInterval := time.Second * 1

	i := 0
	for {
		fmt.Printf("Writing %d to %s\n", i, fileName)
		if err := os.WriteFile(fileName, []byte(strconv.Itoa(i)), 0644); err != nil {
			panic(err)
		}
		i += 1
		if i == maxIter {
			break
		}
		if i == crashIter {
			panic("Program panic!!")
		}
		time.Sleep(timeInterval)
	}
	fmt.Println("Program ended normally!")
}
