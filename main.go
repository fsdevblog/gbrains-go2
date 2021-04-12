package main

import (
	"gbrains-go2/lesson1"
	"log"
)

func main() {
	// Lesson 1.1
	log.Println("Lesson 1.1 start")

	_, err := lesson1.CallPanicFunc()

	if err != nil {
		log.Printf("Got error: %v\n", err)
	}

	log.Println("Lesson 1.1 complete")

	// Lesson 1.2
	log.Println("Lesson 1.2 start")

	err = lesson1.IterateWithOneMillionFiles()

	if err != nil {
		log.Printf("Error %v\n", err)
	}

	log.Println("Lesson 1.2 complete")
}
