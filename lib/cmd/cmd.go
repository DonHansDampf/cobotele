package cmd

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/donhansdampf/cobotele/lib/db"
	"log"
	"os"
	"time"
)

// ComicItem is a struct which defines a fetched comic.
type ComicItem struct {
	Title      string
	PictureURL string
	SiteName   string
	Date       time.Time
}

// ComicSiteTraits defines a comic-site. The number of
// fetched comics in one rune is declared in ComicNum.
// This is importand for calculating the channel size.
type ComicSiteTraits struct {
	SiteName string
	SiteURL  string
	ComicNum int
}

// Bucket gets or creataes a bucket in db.
func (comicSite *ComicSiteTraits) Bucket() *bolt.Bucket {
	bucket, err := db.GetComicSiteBucket(comicSite.SiteName)
	CatchError(err)
	return bucket
}

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

// CatchError prints out the given error.
func CatchError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
