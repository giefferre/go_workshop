package main

import (
	"fmt"
	"log"
	"errors"
	"encoding/json"
	"unicode/utf8"
)


type Person struct {
	FirstName string `json:"personFirstName"`
	LastName string `json:"personLastName"`
	Age int `json:"personAge"`
	bankPin int `json:"personBankPin"`
}

func NewPerson(first string, last string, age int, pin int) (Person, error) {
	var p Person

	if utf8.RuneCountInString(first) > 0 {
		p.FirstName = first
	} else {
		return p, errors.New("First Name can't be blank")
	}

	if utf8.RuneCountInString(last) > 0 {
		p.LastName = last
	} else {
		return p, errors.New("Last Name can't be blank")
	}

	if age > 0 {
		p.Age = age
	} else {
		return p, errors.New("Invalid value for age")
	}

	p.bankPin = pin

	return p, nil
}

func (person *Person) Yell(iterations int) {
	for i := 0; i < iterations; i++ {
		fmt.Println("Zack Yak Yakketi Yak!")
	}
}


func main() {
	// creating a new Person object
	mark, err := NewPerson("Mark", "Zuckemberg", 29, 0000)

	if err != nil {
		log.Fatal(err)
	}

	// let him speak
	mark.Yell(2)


	// encode Person object as json
	jsonData, err := json.Marshal(mark)
	if (err != nil) {
		log.Println(err)
	}

	// print JSON
	// what's missing?
	fmt.Println(string(jsonData))
}