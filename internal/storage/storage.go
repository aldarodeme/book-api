package storage

import (
	"book-api/internal/models"
	"context"
	"github.com/uptrace/bun"
)

type Storage interface {
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	UpdateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	GetBook(ctx context.Context, id int64) (*models.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}

type sto struct {
	db bun.IDB
}

func NewMySqlStorage(db bun.IDB) Storage {
	return &sto{db: db}
}
