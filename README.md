# Book API

### Description

Book API is a RESTful API created using Goa design in Go language.
The API connects to a MySQL database and perform basic CRUD (Create, Read, Update,
Delete) operations on a Book resource.

### Design

The Design package is where the Goa design dsl is written.
Then, the code under gen, which takes care of the http server, marshalling and unmarshalling of requests and responses
is taken care by GOA.
The `Book Service` package is where the business logic lies, and implements the book.Service interface specified by Goa.

### Set up

#### Prerequisites

* Install [Golang](https://go.dev/dl/)
* Install [Docker](https://docs.docker.com/desktop/install/mac-install/)

### Running

* Run `docker compose up` -> This will bring up a mysql database and a UI for it running on localhost:8080. It will also
  bring up the server, which by default, runs on localhost:8000
* Run `go run cmd/migrate/main.go db init` -> This will create the migrations table on the mysql database
* Run `go run cmd/migrate/main.go db migrate` -> This will run the migrations we have on the `migrations` folder for
  creating the `Books` table.

### Test

* After `docker compose up` has finished, and migrations are run, running the command `go test ./...` will run the test
  suite. 
