// Code generated by goa v3.13.2, DO NOT EDIT.
//
// book service
//
// Command:
// $ goa gen book-api/design

package book

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The Book Service service performs multiple operations on book resource.
type Service interface {
	// Create implements Create.
	Create(context.Context, *BookPayload) (res *BookPayload, err error)
	// Update implements Update.
	Update(context.Context, *BookPayload) (res *BookPayload, err error)
	// GetOne implements GetOne.
	GetOne(context.Context, *GetOnePayload) (res *BookPayload, err error)
	// Delete implements Delete.
	Delete(context.Context, *DeletePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "book"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"Create", "Update", "GetOne", "Delete"}

// The book you want to create
type BookPayload struct {
	ID                *int64
	Title             string
	Author            string
	BookCoverImageURL string
	PublishedAt       string
}

// The book you want to delete
type DeletePayload struct {
	// The id of the book you want to delete
	ID *int64
}

// The book you want to get
type GetOnePayload struct {
	// The id of the book you want to get
	ID *int64
}

// MakeCreateBookError builds a goa.ServiceError from an error.
func MakeCreateBookError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Create book error", false, false, false)
}

// MakeUpdateBookError builds a goa.ServiceError from an error.
func MakeUpdateBookError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Update book error", false, false, false)
}

// MakeGetBookError builds a goa.ServiceError from an error.
func MakeGetBookError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Get book error", false, false, false)
}

// MakeDeleteBookError builds a goa.ServiceError from an error.
func MakeDeleteBookError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Delete book error", false, false, false)
}
