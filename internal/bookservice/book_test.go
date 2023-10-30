package bookservice

import (
	"book-api/gen/book"
	"book-api/internal/config"
	"book-api/internal/dependencies"
	"book-api/internal/storage"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestBookScv_Create(t *testing.T) {
	id := int64(1)
	cases := []struct {
		name          string
		payload       *book.BookPayload
		expected      *book.BookPayload
		expectedError error
	}{
		{
			name: "create book success",
			payload: &book.BookPayload{
				Title:             "Test 1",
				Author:            "Test Author",
				BookCoverImageURL: "Test URL",
				PublishedAt:       "2022-01-01T07:08:56Z",
			},
			expected: &book.BookPayload{
				ID:                &id,
				Title:             "Test 1",
				Author:            "Test Author",
				BookCoverImageURL: "Test URL",
				PublishedAt:       "2022-01-01 07:08:56 +0000 UTC",
			},
		},
		{
			name: "create book error no title",
			payload: &book.BookPayload{
				Author:            "Test Author",
				BookCoverImageURL: "Test URL",
				PublishedAt:       "2022-01-01T07:08:56Z",
			},
			expectedError: errors.New("title can not be empty"),
		},
	}
	cfg, err := config.LoadConfig("../../.env")
	require.NoError(t, err)
	deps, err := dependencies.InitDependencies(cfg)
	require.NoError(t, err)
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tx, err := deps.DB.Begin()
			require.NoError(t, err)
			sto := storage.NewMySqlStorage(tx)
			s := NewBookService(log.Default(), sto)

			created, err := s.Create(context.Background(), c.payload)
			if c.expectedError != nil {
				assert.Equal(t, c.expectedError, err)
				assert.Nil(t, created)
				return
			}
			assert.NoError(t, err)
			// Force equal ID as it is auto incremented
			c.expected.ID = created.ID
			assert.Equal(t, c.expected, created)
			err = tx.Rollback()
			require.NoError(t, err)
		})
	}
}
