package cobo

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/donhansdampf/cobotele/lib/cmd"
	"github.com/donhansdampf/cobotele/lib/db"
	"github.com/donhansdampf/cobotele/lib/poorlydrawnlines"
)

// Start launches the fetching process.
func Start() {
	comicSitesList := createComicSiteList()
	comicNumSum := sumComicNum(comicSitesList)

	comicQueue := make(chan *cmd.ComicItem, comicNumSum)

	go poorlydrawnlines.GetComic(comicQueue)

	for i := 0; i < comicNumSum; i++ {
		comic := <-comicQueue
		logMsg := fmt.Sprintf("Fetched Comic: %s", comic)
		cmd.PrintVerbose(logMsg)

		handleComicItem(comic)

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

func handleComicItem(comicItem *cmd.ComicItem) {
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
	// Send to telegramGroup.
	fmt.Println(comicItemBucket)
	return
}
