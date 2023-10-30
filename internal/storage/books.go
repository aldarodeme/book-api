package storage

import (
	"book-api/internal/models"
	"context"
)

func (s *sto) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	res, err := s.db.NewInsert().Model(book).Exec(ctx)
	if err != nil {
		return nil, err
	}
	insertedID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.Id = insertedID

	return book, nil
}

func (s *sto) UpdateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	_, err := s.db.NewUpdate().Model(book).Where("id = ?", book.Id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *sto) GetBook(ctx context.Context, id int64) (*models.Book, error) {
	var res models.Book

	_, err := s.db.NewSelect().Model(&res).Where("id = ?", id).Exec(ctx, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *sto) DeleteBook(ctx context.Context, id int64) error {
	_, err := s.db.NewDelete().Model(&models.Book{}).Where("id = ?", id).Exec(ctx)
	return err
}
