package main

import (
	"log"
	"flag"
	"fmt"
	"time"
	"math/rand"
)

var UserChoice string

func init() {
	flag.StringVar(&UserChoice, "example", "first", "Example code to run")
}

func main() {
	fmt.Println("\nHello from this multi example program\n\n")

	// this binds the real values of the input flags, if present
	flag.Parse()

	switch UserChoice {
	case "first":
		runGoroutines()

	case "second":
		runChannels()

	case "third":
		runSelect()

	default:
		log.Fatal("Invalid example")
	}
}




func yell(word string, seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println(word)
}

func runGoroutines() {
	go yell("On", 3)
	go yell("My", 1)
	go yell("Way", 2)

	time.Sleep(time.Duration(5) * time.Second)
}




func channelWriter(aChan chan int) {
	rand.Seed( time.Now().UTC().UnixNano() )

	time.Sleep(time.Duration(1) * time.Second)
	aChan <- rand.Intn(10)
}

func runChannels() {
	myChan := make(chan int)

	go channelWriter(myChan)

	fromChan := <-myChan
	fmt.Println("Received ", fromChan)
}




func runSelect() {
	// both time.After and time.NewTicker return a channel
	tick := time.NewTicker(500 * time.Millisecond)
	boom := time.After(4 * time.Second)

	// the bomb!
	bombLoop:
	for {
		select {
			case <-tick.C:
				fmt.Println("tick!")

			case <-boom:
				fmt.Println("BOOOOOM!")
				break bombLoop
		}
	}
}