// Code generated by goa v3.5.2, DO NOT EDIT.
//
// exmp HTTP client CLI support package
//
// Command:
// $ goa gen github.com/lockwobr/goa-example-multi-type/pkg/design -o pkg

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	exmpc "github.com/lockwobr/goa-example-multi-type/pkg/gen/http/exmp/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `exmp (get-person|get-fam)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` exmp get-person` + "\n" +
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
) (goa.Endpoint, interface{}, error) {
	var (
		exmpFlags = flag.NewFlagSet("exmp", flag.ContinueOnError)

		exmpGetPersonFlags = flag.NewFlagSet("get-person", flag.ExitOnError)

		exmpGetFamFlags = flag.NewFlagSet("get-fam", flag.ExitOnError)
	)
	exmpFlags.Usage = exmpUsage
	exmpGetPersonFlags.Usage = exmpGetPersonUsage
	exmpGetFamFlags.Usage = exmpGetFamUsage

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
		case "exmp":
			svcf = exmpFlags
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
		case "exmp":
			switch epn {
			case "get-person":
				epf = exmpGetPersonFlags

			case "get-fam":
				epf = exmpGetFamFlags

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
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "exmp":
			c := exmpc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-person":
				endpoint = c.GetPerson()
				data = nil
			case "get-fam":
				endpoint = c.GetFam()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// exmpUsage displays the usage of the exmp command and its subcommands.
func exmpUsage() {
	fmt.Fprintf(os.Stderr, `Service is the exmp service interface.
Usage:
    %[1]s [globalflags] exmp COMMAND [flags]

COMMAND:
    get-person: GetPerson implements get-person.
    get-fam: GetFam implements get-fam.

Additional help:
    %[1]s exmp COMMAND --help
`, os.Args[0])
}
func exmpGetPersonUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] exmp get-person

GetPerson implements get-person.

Example:
    %[1]s exmp get-person
`, os.Args[0])
}

func exmpGetFamUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] exmp get-fam

GetFam implements get-fam.

Example:
    %[1]s exmp get-fam
`, os.Args[0])
}
