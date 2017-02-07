package main

import (
	"fmt"
	"github.com/donhansdampf/cobotele/lib/cmd"
	"log"
)

func main() {
	telegramToken := cmd.GetFlags()
	if telegramToken == "none" {
		log.Fatalln("No token provided." +
			"Please use one via 'cobotele --token=TOKEN'.")
	}

	fmt.Println("Test")
}
