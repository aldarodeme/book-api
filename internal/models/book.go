package models

import (
	"errors"
	"time"
)

type Book struct {
	Id                int64     `json:"id"`
	Title             string    `json:"title"`
	Author            string    `json:"author"`
	BookCoverImageURL string    `json:"book_cover_image_url"`
	PublishedAt       time.Time `json:"published_at"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("title can not be empty")
	}

	if b.Author == "" {
		return errors.New("author can not be empty")
	}

	if b.BookCoverImageURL == "" {
		return errors.New("BookCoverImageURL can not be empty")
	}

	return nil
}
