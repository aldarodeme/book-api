// Code generated by goa v3.13.2, DO NOT EDIT.
//
// book HTTP client CLI support package
//
// Command:
// $ goa gen book-api/design

package client

import (
	book "book-api/gen/book"
	"encoding/json"
	"fmt"
	"strconv"
)

// BuildCreatePayload builds the payload for the book Create endpoint from CLI
// flags.
func BuildCreatePayload(bookCreateBody string) (*book.BookPayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(bookCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"Ut est sit voluptatem.\",\n      \"book_cover_image_url\": \"Deserunt ipsum non dignissimos in officia.\",\n      \"id\": 4314128948732407298,\n      \"published_at\": \"Soluta in provident veritatis natus.\",\n      \"title\": \"Nulla porro officiis.\"\n   }'")
		}
	}
	v := &book.BookPayload{
		ID:                body.ID,
		Title:             body.Title,
		Author:            body.Author,
		BookCoverImageURL: body.BookCoverImageURL,
		PublishedAt:       body.PublishedAt,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the book Update endpoint from CLI
// flags.
func BuildUpdatePayload(bookUpdateBody string, bookUpdateID string) (*book.BookPayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(bookUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"In architecto nisi magnam et sed autem.\",\n      \"book_cover_image_url\": \"Et suscipit dolorem amet a beatae iusto.\",\n      \"published_at\": \"Aut sunt id eum libero qui.\",\n      \"title\": \"Soluta ut enim modi ex magnam.\"\n   }'")
		}
	}
	var id int64
	{
		id, err = strconv.ParseInt(bookUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	v := &book.BookPayload{
		Title:             body.Title,
		Author:            body.Author,
		BookCoverImageURL: body.BookCoverImageURL,
		PublishedAt:       body.PublishedAt,
	}
	v.ID = &id

	return v, nil
}

// BuildGetOnePayload builds the payload for the book GetOne endpoint from CLI
// flags.
func BuildGetOnePayload(bookGetOneID string) (*book.GetOnePayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(bookGetOneID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	v := &book.GetOnePayload{}
	v.ID = &id

	return v, nil
}

// BuildDeletePayload builds the payload for the book Delete endpoint from CLI
// flags.
func BuildDeletePayload(bookDeleteID string) (*book.DeletePayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(bookDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	v := &book.DeletePayload{}
	v.ID = &id

	return v, nil
}
