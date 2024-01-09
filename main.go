package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var maxIter int
	var crashIter int
	var fileName string
	flag.IntVar(&maxIter, "runTime", -1, "Makes program exit ok after this amount of seconds. -1 (default) means infinite.")
	flag.IntVar(&crashIter, "crashTime", -1, "Makes program crash after this amount if seconds. -1 (default) means infinite.")
	flag.StringVar(&fileName, "fName", "./output.txt", "Filename for the output.")
	flag.Parse()

	fmt.Println("Starting program")

	timeInterval := time.Second * 1

	// Setup listerner for SIGTERM and SIGINT (ctrl+c)
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	// Setup listener for goroutine exit
	done := make(chan bool)

	os.Truncate(fileName, 0)

	go func() {
		i := 0
		for {
			fmt.Printf("Writing %d to %s\n", i, fileName)
			if err := writeToFile(fileName, fmt.Sprintf("%d\n", i)); err != nil {
				panic(err)
			}
			if i == maxIter {
				writeToFile(fileName, "Program reached max runtime.\n")
				break
			}
			if i == crashIter {
				writeToFile(fileName, "Program panic!!\n")
				panic("Program panic!!")
			}
			time.Sleep(timeInterval)
			i += 1
		}
		done <- true
	}()

loop:
	for {
		select {
		case sig := <-cancelChan:
			fmt.Printf("Caught signal %q\n", sig)
			writeToFile(fileName, fmt.Sprintf("Caught signal %v\n", sig))
			break loop
		case <-done:
			break loop
		}
	}

	fmt.Println("Program ended normally!")
	writeToFile(fileName, "Program ended normally!")
}

func writeToFile(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err == nil {
		defer f.Close()

		_, err = f.WriteString(text)
	}
	return err
}
