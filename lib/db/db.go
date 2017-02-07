package db

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

// GetComicSiteBucket creates a bucket for the given comic-site.
func GetComicSiteBucket(bucketName string) (*bolt.Bucket, error) {
	comicDatabase, err := bolt.Open("comics.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer comicDatabase.Close()

	var bucket *bolt.Bucket
	err = comicDatabase.Update(func(tx *bolt.Tx) error {
		bucket, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})

	return bucket, err
}

// GetComicItemBucket creates a bucket. Returns ErrBucketExists if comic already in db.
func GetComicItemBucket(comicItemTitle string, comicItemSiteName string) (*bolt.Bucket, error) {
	comicDatabase, err := bolt.Open("comics.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer comicDatabase.Close()

	tx, err := comicDatabase.Begin(true)
	defer tx.Rollback()

	comicSiteBucket := tx.Bucket([]byte(comicItemSiteName))

	comicItemBucket, err := comicSiteBucket.CreateBucket([]byte(comicItemTitle))
	if err != nil {
		return comicItemBucket, err
	}

	if err := tx.Commit(); err != nil {
		return comicItemBucket, err
	}

	return comicItemBucket, nil
}

// InsertComicItem puts Traits of an ComicItem into its bucket.
func InsertComicItem(comicItemBucket *bolt.Bucket, title string, pictureURL string, siteName string, date time.Time) error {
	comicDatabase, err := bolt.Open("comics.db", 0600, nil)
	if err != nil {
		return err
	}
	defer comicDatabase.Close()

	comicDatabase.Update(func(tx *bolt.Tx) error {
		dateStr := date.Format(time.UnixDate)

		err = comicItemBucket.Put([]byte("Title"), []byte(title))
		err = comicItemBucket.Put([]byte("PictureURL"), []byte(pictureURL))
		err = comicItemBucket.Put([]byte("SiteName"), []byte(siteName))
		err = comicItemBucket.Put([]byte("Date"), []byte(dateStr))
		return err
	})

	return nil
}
