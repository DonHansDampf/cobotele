package cobo

import "testing"

func TestSumComicNum(t *testing.T) {
	comicSiteList := createComicSiteList()
	comicNumSum := sumComicNum(comicSiteList)

	if comicNumSum < 1 {
		t.Error("comicNumSum no number or below 1.")
	}
}
