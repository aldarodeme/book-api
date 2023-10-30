package design

import (
	. "goa.design/goa/v3/dsl"
)

var BookPayload = Type("BookPayload", func() {
	Attribute("id", Int64)
	Attribute("title", String)
	Attribute("author", String)
	Attribute("book_cover_image_url", String)
	Attribute("published_at", String)

	Required("title", "author", "book_cover_image_url", "published_at")

})

var _ = API("book", func() {
	Title("Book API service")
	Description("Service for managing books")
	Server("book-api", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("book", func() {
	Description("The Book Service service performs multiple operations on book resource.")

	Method("Create", func() {
		Payload(BookPayload, "The book you want to create")

		// Error defines an error result.
		Error("Create book error")

		Result(BookPayload, "The created book")

		HTTP(func() {
			POST("/books")
		})
	})

	Method("Update", func() {
		Payload(BookPayload, "The book you want to update")

		// Error defines an error result.
		Error("Update book error")

		Result(BookPayload, "The updated book")

		HTTP(func() {
			PUT("/books/{id}")
		})
	})

	Method("GetOne", func() {
		Payload(func() {
			Attribute("id", Int64, "The id of the book you want to get")
		}, "The book you want to get")

		// Error defines an error result.
		Error("Get book error")

		Result(BookPayload, "The book requested")

		HTTP(func() {
			GET("/books/{id}")
		})
	})

	Method("Delete", func() {
		Payload(func() {
			Attribute("id", Int64, "The id of the book you want to delete")
		}, "The book you want to delete")

		// Error defines an error result.
		Error("Delete book error")

		HTTP(func() {
			DELETE("/books/{id}")
		})
	})

	Files("/openapi.json", "./gen/http.go/openapi.json")
})
