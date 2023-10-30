package bookservice

import (
	"book-api/internal/models"
	"book-api/internal/storage"
	"context"
	"log"
	"time"

	"book-api/gen/book"
)

// book service example implementation.
// The example methods log the requests and return zero values.
type bookScv struct {
	logger *log.Logger
	sto    storage.Storage
}

// Create will take care of creating a book.
func (s *bookScv) Create(ctx context.Context, payload *book.BookPayload) (*book.BookPayload, error) {
	publishedAt, err := time.Parse(time.RFC3339, payload.PublishedAt)
	if err != nil {
		return nil, err
	}

	toCreateBook := &models.Book{
		Title:             payload.Title,
		Author:            payload.Author,
		BookCoverImageURL: payload.BookCoverImageURL,
		PublishedAt:       publishedAt,
	}

	err = toCreateBook.Validate()
	if err != nil {
		return nil, err
	}

	newBook, err := s.sto.CreateBook(ctx, toCreateBook)
	if err != nil {
		return nil, book.MakeCreateBookError(err)
	}

	return &book.BookPayload{
		ID:                &newBook.Id,
		Title:             newBook.Title,
		Author:            newBook.Author,
		BookCoverImageURL: newBook.BookCoverImageURL,
		PublishedAt:       newBook.PublishedAt.String(),
	}, nil
}

func (s *bookScv) Update(ctx context.Context, payload *book.BookPayload) (res *book.BookPayload, err error) {
	publishedAt, err := time.Parse(time.RFC3339, payload.PublishedAt)
	if err != nil {
		return nil, err
	}
	updatedBook, err := s.sto.UpdateBook(ctx, &models.Book{
		Id:                *payload.ID,
		Title:             payload.Title,
		Author:            payload.Author,
		BookCoverImageURL: payload.BookCoverImageURL,
		PublishedAt:       publishedAt,
	})
	if err != nil {
		return nil, err
	}

	return &book.BookPayload{
		ID:                &updatedBook.Id,
		Title:             updatedBook.Title,
		Author:            updatedBook.Author,
		BookCoverImageURL: updatedBook.BookCoverImageURL,
		PublishedAt:       updatedBook.PublishedAt.String(),
	}, nil
}

func (s *bookScv) GetOne(ctx context.Context, payload *book.GetOnePayload) (res *book.BookPayload, err error) {
	bookRes, err := s.sto.GetBook(ctx, *payload.ID)
	if err != nil {
		return nil, err
	}

	return &book.BookPayload{
		ID:                &bookRes.Id,
		Title:             bookRes.Title,
		Author:            bookRes.Author,
		BookCoverImageURL: bookRes.BookCoverImageURL,
		PublishedAt:       bookRes.PublishedAt.String(),
	}, nil
}

func (s *bookScv) Delete(ctx context.Context, payload *book.DeletePayload) (err error) {
	return s.sto.DeleteBook(ctx, *payload.ID)
}

// NewBookService returns the calc service implementation.
func NewBookService(logger *log.Logger, sto storage.Storage) book.Service {
	return &bookScv{
		logger: logger,
		sto:    sto,
	}
}
