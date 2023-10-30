// Code generated by goa v3.13.2, DO NOT EDIT.
//
// book HTTP server
//
// Command:
// $ goa gen book-api/design

package server

import (
	book "book-api/gen/book"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the book service endpoint HTTP handlers.
type Server struct {
	Mounts               []*MountPoint
	Create               http.Handler
	Update               http.Handler
	GetOne               http.Handler
	Delete               http.Handler
	GenHTTPGoOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the book service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *book.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemGenHTTPGoOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPGoOpenapiJSON == nil {
		fileSystemGenHTTPGoOpenapiJSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Create", "POST", "/books"},
			{"Update", "PUT", "/books/{id}"},
			{"GetOne", "GET", "/books/{id}"},
			{"Delete", "DELETE", "/books/{id}"},
			{"./gen/http.go/openapi.json", "GET", "/openapi.json"},
		},
		Create:               NewCreateHandler(e.Create, mux, decoder, encoder, errhandler, formatter),
		Update:               NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		GetOne:               NewGetOneHandler(e.GetOne, mux, decoder, encoder, errhandler, formatter),
		Delete:               NewDeleteHandler(e.Delete, mux, decoder, encoder, errhandler, formatter),
		GenHTTPGoOpenapiJSON: http.FileServer(fileSystemGenHTTPGoOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "book" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Create = m(s.Create)
	s.Update = m(s.Update)
	s.GetOne = m(s.GetOne)
	s.Delete = m(s.Delete)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return book.MethodNames[:] }

// Mount configures the mux to serve the book endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCreateHandler(mux, h.Create)
	MountUpdateHandler(mux, h.Update)
	MountGetOneHandler(mux, h.GetOne)
	MountDeleteHandler(mux, h.Delete)
	MountGenHTTPGoOpenapiJSON(mux, goahttp.Replace("", "/./gen/http.go/openapi.json", h.GenHTTPGoOpenapiJSON))
}

// Mount configures the mux to serve the book endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountCreateHandler configures the mux to serve the "book" service "Create"
// endpoint.
func MountCreateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/books", f)
}

// NewCreateHandler creates a HTTP handler which loads the HTTP request and
// calls the "book" service "Create" endpoint.
func NewCreateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateRequest(mux, decoder)
		encodeResponse = EncodeCreateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "book")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateHandler configures the mux to serve the "book" service "Update"
// endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/books/{id}", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "book" service "Update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "book")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetOneHandler configures the mux to serve the "book" service "GetOne"
// endpoint.
func MountGetOneHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/books/{id}", f)
}

// NewGetOneHandler creates a HTTP handler which loads the HTTP request and
// calls the "book" service "GetOne" endpoint.
func NewGetOneHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetOneRequest(mux, decoder)
		encodeResponse = EncodeGetOneResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "GetOne")
		ctx = context.WithValue(ctx, goa.ServiceKey, "book")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteHandler configures the mux to serve the "book" service "Delete"
// endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/books/{id}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "book" service "Delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, decoder)
		encodeResponse = EncodeDeleteResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "book")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPGoOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPGoOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}