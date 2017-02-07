package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var verbose bool

// GetFlags returns the Telegram-Token provided by the flag "token"
func GetFlags() string {
	telegramToken := flag.String("token", "none", "Token for Telegram-Bot")
	helpBool := flag.Bool("h", false, "Show help.")
	flag.BoolVar(&verbose, "v", false, "Enable verbose Logging.")

	flag.Parse()

	if *helpBool == true {
		printHelpText()
	}

	return *telegramToken
}

func printHelpText() {
	fmt.Printf("%s\n\n%s\n\t%s\n\n%s\n\t%s\t%s\n\t%s\t%s\n\t%s\t%s\n",
		"Cobotele is a comic-Bot for Telegram-Grops",
		"Usage:",
		"cobotele [arguments]",
		"Flags",
		"token",
		"provides token for telegram",
		"v",
		"enables verbose logging",
		"h",
		"shows this help",
	)
	os.Exit(2)
}

// PrintVerbose prints out log msg if flag "verbose" is set.
func PrintVerbose(msg string) {
	if verbose == true {
		log.Println(msg)
	}
}
