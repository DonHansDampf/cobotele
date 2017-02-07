package cobo

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/donhansdampf/cobotele/lib/cmd"
	"github.com/donhansdampf/cobotele/lib/db"
	"github.com/donhansdampf/cobotele/lib/poorlydrawnlines"
	"github.com/donhansdampf/cobotele/lib/telegram"
	"github.com/tucnak/telebot"
)

// Start launches the fetching process.
func Start(telegramToken string) {
	bot, _ := telegram.New(telegramToken)

	comicSitesList := createComicSiteList()
	comicNumSum := sumComicNum(comicSitesList)

	comicQueue := make(chan *cmd.ComicItem, comicNumSum)

	go poorlydrawnlines.GetComic(comicQueue)

	for i := 0; i < comicNumSum; i++ {
		comic := <-comicQueue
		logMsg := fmt.Sprintf("Fetched Comic: %s", comic)
		cmd.PrintVerbose(logMsg)

		handleComicItem(comic, bot)

	}
}

func createComicSiteList() []*cmd.ComicSiteTraits {
	comicSiteList := []*cmd.ComicSiteTraits{}

	poorlyDrawnTraits := &cmd.ComicSiteTraits{
		SiteName: "PoorlyDrawnLines",
		SiteURL:  "http://poorlydrawnlines.com",
		ComicNum: 10,
	}
	comicSiteList = append(comicSiteList, poorlyDrawnTraits)
	_ = poorlyDrawnTraits.Bucket()

	// Repeat for other comic-Sites. Way to automate this better?

	for _, comicSite := range comicSiteList {
		logMsg := fmt.Sprintf("Added comic %s to list.",
			comicSite.SiteName)
		cmd.PrintVerbose(logMsg)
	}

	return comicSiteList
}

func sumComicNum(comicSitesTraits []*cmd.ComicSiteTraits) int {
	var comicNumSum int

	for _, comicSiteTraits := range comicSitesTraits {
		comicNumSum += comicSiteTraits.ComicNum
	}

	return comicNumSum
}

func handleComicItem(comicItem *cmd.ComicItem, bot *telebot.Bot) {
	comicItemBucket, err := db.GetComicItemBucket(comicItem.Title, comicItem.SiteName)

	if err == bolt.ErrBucketExists {
		logMsg := fmt.Sprintf("Bucket for Comic '%s' already exists. Skipped.",
			comicItem.Title)
		cmd.PrintVerbose(logMsg)
		return
	}

	// Insert Informations into comicItemBucket.
	err = db.InsertComicItem(
		comicItemBucket,
		comicItem.Title,
		comicItem.PictureURL,
		comicItem.SiteName,
		comicItem.Date,
	)
	cmd.CatchError(err)

	// Download pictureSRC to disk.
	filepath, err := db.DownloadComicPicture(comicItem.PictureURL, comicItem.SiteName)
	cmd.CatchError(err)

	// Send Comic to Telegram Group.
	err = telegram.SendComic(bot, 88888, comicItem.Title, filepath)
	cmd.CatchError(err)
	return
}
