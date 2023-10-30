FROM golang:1.21.1-alpine3.18 AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update

# Create appuser.
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR app

COPY . .

# Build the binary and make it executable.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/book_api ./cmd/book-api/main.go ./cmd/book-api/http.go

RUN chmod +x /go/bin/book_api

FROM scratch
# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable.
COPY --from=builder /go/bin/book_api /go/bin/book_api
# Use an unprivileged user.
USER appuser:appuser

# Run the book_api binary.
ENTRYPOINT ["/go/bin/book_api"]