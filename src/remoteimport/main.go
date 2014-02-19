package main

import (
	"github.com/gdlabs/goexamples/hellogo"
)

func main() {
	hellogo.SayHello()
	hellogo.yellRandomly()
}


/* for convenience the content of the hellogo package is hereafter shown

package hellogo

import "fmt" 

func SayHello() {
        fmt.Println("Hello World from hellogo package!")
}

func yellRandomly() {
        fmt.Println("Zack Yak Yakketi Yak!")
}

 */