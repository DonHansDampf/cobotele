package main

import (
	"fmt"
	"github.com/donhansdampf/cobotele/lib"
	"log"
)

func main() {
	telegramToken, verboseBool := cmd.GetFlags()
	if telegramToken == "none" {
		log.Fatalln("No token provided. Please use one via 'cobotele --token=TOKEN'.")
	}

	fmt.Println("Token:", telegramToken)
	fmt.Println("Verbose:", verboseBool)
}
