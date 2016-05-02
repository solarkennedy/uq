package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
)

func readData(format string, filename string) []byte {
	var bytes []byte
	if filename == "-" {
		bytes, _ = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, _ = ioutil.ReadFile(os.Args[len(os.Args)-1])
	}
	return bytes
}

func parseData(input_data []byte, format string) (parsed_data interface{}, err error) {
	// TODO: Support more formats :)
	err = yaml.Unmarshal(input_data, &parsed_data)
	return parsed_data, err
}

func outputData(input_data interface{}, format string) (output_data []byte, err error) {
	// TODO: Support more formats
	if format == "json" {
		output_data, err = json.MarshalIndent(input_data, "", "  ")
		// TODO: Detect auto better
	} else if format == "yaml" || format == "yml" || format == "auto" {
		output_data, err = yaml.Marshal(input_data)
	} else {
		fmt.Println(string(json))
		return 0
	}
}

func parseArgs() map[string]interface{} {
	usage := `uq - Universal (de)serializer

Usage:
  uq [-s FORMAT] [-t FORMAT] [FILE]

Options:
  -s <FORMAT>, --source <FORMAT>  Specify input format. [default: auto]
  -t <FORMAT>, --target <FORMAT>  Specify output format. [default: auto]
  -h --help     Show this screen
  --version     Show version

Formats:
  * json
  * yaml|yml
`

	arguments, _ := docopt.Parse(usage, nil, true, "0.0.1", false)
	return arguments
}

func main() {
	args := parseArgs()
	fmt.Print("args: ")
	fmt.Println(args)
	var filename string
	if args["FILE"] == nil {
		filename = "-"
	} else {
		filename = args["FILE"].(string)
	}
	data := readData(args["--source"].(string), filename)
	exitcode := outputData(data)
	os.Exit(exitcode)
}
