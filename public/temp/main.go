package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	err := fmt.Errorf("error occurredered")
	fmt.Println("Hello, World!")
	if err != nil {
		fmt.Println("error occurredered")
		log.Fatal(err)
		fmt.Println("error occurredered")
	}
}
