package cobo

import (
	"fmt"
	"github.com/donhansdampf/cobotele/lib/cmd"
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
