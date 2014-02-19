package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
	"basics/utilities"
)


func main() {
	/* PART 1
	 * variables and if conditions
	 */

	fmt.Println("--- PART 1 ---")

	var myStringVar string // new instance of a string

	myIntvar := 1 // new instance of an integer

	//myStringVar := "Hello Mum"	// this would generate an error!
	myStringVar = "Hello Dad"		// it is ok


	/*
	if myIntvar == 2				// THIS WILL NOW WORK
	{
		someCode()
	}
	*/

	if myIntvar == 1 {				// bracketing ok

		_ = utilities.SomeCode()	// what's this?

	}

	fmt.Println(myStringVar)
	fmt.Println(myIntvar)





	/* PART 2
	 * error handling
	 */

	fmt.Println("\n\n--- PART 2 ---")

	// there is no exception, we HAVE to check for errors

	filePointer, error := os.Open("file.txt")
	if error != nil {
		fmt.Println("Error handling here")
	}

	// if everything is fine we are considering no error is present
	// in this case it's ok to proceed, but... what if something went bad?

	filePointer.Close()





	/* PART 3
	 * switch, random, time
	 */

	fmt.Println("\n\n--- PART 3 ---")

	rand.Seed( time.Now().UTC().UnixNano() )
	aRandomNumber := rand.Intn(10)

	fmt.Print( fmt.Sprintf("%d is ", aRandomNumber) )

	switch {
	case aRandomNumber < 5:
		fmt.Println(" lower than 5")

	case aRandomNumber >= 5:
		fmt.Println(" greater than or equal to 5")
	}


	time.Sleep(1 * time.Second)
	fmt.Println("Quit...")
}