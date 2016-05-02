package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
)

func readData(filename string) (bytes []byte, err error) {
	if filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(os.Args[len(os.Args)-1])
	}
	return bytes, err
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
		panic("Non supported output format")
	}

	return output_data, err
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
	data, _ := readData(filename)

	parsed_data, _ := parseData(data, args["--source"].(string))

	output_data, _ := outputData(parsed_data, args["--target"].(string))

	fmt.Println(string(output_data))
}
