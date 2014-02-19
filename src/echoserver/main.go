package main

import (
	"io"
	"log"
	"net"
)

const address = "localhost:8000"

func main() {
	listenObj, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening on", address)

	for {
		connection, err := listenObj.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("New incoming connection!")

		go io.Copy(connection, connection)
	}
}