package main

import (
	"github.com/donhansdampf/cobotele/lib/cmd"
	"github.com/donhansdampf/cobotele/lib/cobo"
	"log"
)

func main() {
	telegramToken := cmd.GetFlags()
	if telegramToken == "none" {
		log.Fatalln("No token provided." +
			"Please use one via 'cobotele --token=TOKEN'.")
	}

	cobo.Start()
}
