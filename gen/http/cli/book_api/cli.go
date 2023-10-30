// Code generated by goa v3.13.2, DO NOT EDIT.
//
// book-api HTTP client CLI support package
//
// Command:
// $ goa gen book-api/design

package cli

import (
	bookc "book-api/gen/http/book/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `book (create|update|get-one|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` book create --body '{
      "author": "Ut est sit voluptatem.",
      "book_cover_image_url": "Deserunt ipsum non dignissimos in officia.",
      "id": 4314128948732407298,
      "published_at": "Soluta in provident veritatis natus.",
      "title": "Nulla porro officiis."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		bookFlags = flag.NewFlagSet("book", flag.ContinueOnError)

		bookCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		bookCreateBodyFlag = bookCreateFlags.String("body", "REQUIRED", "")

		bookUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		bookUpdateBodyFlag = bookUpdateFlags.String("body", "REQUIRED", "")
		bookUpdateIDFlag   = bookUpdateFlags.String("id", "REQUIRED", "")

		bookGetOneFlags  = flag.NewFlagSet("get-one", flag.ExitOnError)
		bookGetOneIDFlag = bookGetOneFlags.String("id", "REQUIRED", "The id of the book you want to get")

		bookDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		bookDeleteIDFlag = bookDeleteFlags.String("id", "REQUIRED", "The id of the book you want to delete")
	)
	bookFlags.Usage = bookUsage
	bookCreateFlags.Usage = bookCreateUsage
	bookUpdateFlags.Usage = bookUpdateUsage
	bookGetOneFlags.Usage = bookGetOneUsage
	bookDeleteFlags.Usage = bookDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "book":
			svcf = bookFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "book":
			switch epn {
			case "create":
				epf = bookCreateFlags

			case "update":
				epf = bookUpdateFlags

			case "get-one":
				epf = bookGetOneFlags

			case "delete":
				epf = bookDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "book":
			c := bookc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = bookc.BuildCreatePayload(*bookCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = bookc.BuildUpdatePayload(*bookUpdateBodyFlag, *bookUpdateIDFlag)
			case "get-one":
				endpoint = c.GetOne()
				data, err = bookc.BuildGetOnePayload(*bookGetOneIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = bookc.BuildDeletePayload(*bookDeleteIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// bookUsage displays the usage of the book command and its subcommands.
func bookUsage() {
	fmt.Fprintf(os.Stderr, `The Book Service service performs multiple operations on book resource.
Usage:
    %[1]s [globalflags] book COMMAND [flags]

COMMAND:
    create: Create implements Create.
    update: Update implements Update.
    get-one: GetOne implements GetOne.
    delete: Delete implements Delete.

Additional help:
    %[1]s book COMMAND --help
`, os.Args[0])
}
func bookCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book create -body JSON

Create implements Create.
    -body JSON: 

Example:
    %[1]s book create --body '{
      "author": "Ut est sit voluptatem.",
      "book_cover_image_url": "Deserunt ipsum non dignissimos in officia.",
      "id": 4314128948732407298,
      "published_at": "Soluta in provident veritatis natus.",
      "title": "Nulla porro officiis."
   }'
`, os.Args[0])
}

func bookUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book update -body JSON -id INT64

Update implements Update.
    -body JSON: 
    -id INT64: 

Example:
    %[1]s book update --body '{
      "author": "In architecto nisi magnam et sed autem.",
      "book_cover_image_url": "Et suscipit dolorem amet a beatae iusto.",
      "published_at": "Aut sunt id eum libero qui.",
      "title": "Soluta ut enim modi ex magnam."
   }' --id 4371083667502014282
`, os.Args[0])
}

func bookGetOneUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book get-one -id INT64

GetOne implements GetOne.
    -id INT64: The id of the book you want to get

Example:
    %[1]s book get-one --id 5555553311130906412
`, os.Args[0])
}

func bookDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book delete -id INT64

Delete implements Delete.
    -id INT64: The id of the book you want to delete

Example:
    %[1]s book delete --id 7243459915293518351
`, os.Args[0])
}