package telegram

import (
	"github.com/tucnak/telebot"
)

// New creates new instance of a telegram-bot with
// the given token.
func New(telegramToken string) (*telebot.Bot, error) {
	bot, err := telebot.NewBot(telegramToken)
	return bot, err
}

// Sends comic to given Chat-ID.
func SendComic(bot *telebot.Bot, id int64, comicTitle string, filepath string) error {
	chatGroup := telebot.Chat{
		ID:        id,
		Type:      "supergroup",
		Title:     "",
		FirstName: "",
		LastName:  "",
		Username:  "",
	}

	comicPicture, err := createPhoto(filepath, comicTitle)
	if err != nil {
		return err
	}
	bot.SendPhoto(chatGroup, comicPicture, nil)

	return nil
}

func createPhoto(filepath string, comicTitle string) (*telebot.Photo, error) {
	comicPictureFile, err := telebot.NewFile(filepath)
	if err != nil {
		return nil, err
	}

	// Use Picture also as Thumbnail
	comicPictureThumb := telebot.Thumbnail{
		File:   comicPictureFile,
		Width:  200,
		Height: 200,
	}

	comicPicture := &telebot.Photo{
		File:      comicPictureFile,
		Thumbnail: comicPictureThumb,
		Caption:   comicTitle,
	}

	return comicPicture, nil
}
