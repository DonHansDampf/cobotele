package poorlydrawnlines

import (
	"github.com/SlyMarbo/rss"
	"github.com/donhansdampf/cobotele/lib/cmd"
	"regexp"
	"strings"
)

var (
	feedURL      = "https://feeds.feedburner.com/PoorlyDrawnLines"
	pictureRegex = "(http.*\\.png\"\\s)"
)

// GetComic fetches comics via a rss-feed.
func GetComic(comicQueue chan *cmd.ComicItem) {
	feed, err := rss.Fetch(feedURL)
	cmd.CatchError(err)

	for _, item := range feed.Items {
		regex := regexp.MustCompile(pictureRegex)
		pictureSrc := regex.FindStringSubmatch(item.Content)[0]
		pictureSrc = strings.TrimSuffix(pictureSrc, "\" ")

		comicItem := &cmd.ComicItem{
			Title:      item.Title,
			PictureURL: pictureSrc,
			SiteName:   "PoorlyDrawnLines",
			Date:       item.Date,
		}

		comicQueue <- comicItem
	}
}
