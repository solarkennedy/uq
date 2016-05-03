package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	if format == "json" {
		err = json.Unmarshal(input_data, &parsed_data)
	} else if format == "yaml" || format == "yml" {
		err = yaml.Unmarshal(input_data, &parsed_data)
	} else {
		panic("Non support inport format")
	}
	return parsed_data, err
}

func outputData(input_data interface{}, format string) (output_data []byte, err error) {
	if format == "json" {
		output_data, err = json.MarshalIndent(input_data, "", "  ")
	} else if format == "yaml" || format == "yml" {
		output_data, err = yaml.Marshal(input_data)
	} else {
		panic("Non supported output format")
	}

	return output_data, err
}

func detectInputFormat(data []byte) (detected_format string) {
	// TODO: Some sort of input detection
	return "yaml"
}

func parseArgs() map[string]interface{} {
	usage := `uq - Universal (de)serializer

Usage:
  uq [-v] [-s FORMAT] [-t FORMAT] [FILE]

Options:
  -s <FORMAT>, --source <FORMAT>  Specify input format. [default: auto]
  -t <FORMAT>, --target <FORMAT>  Specify output format. [default: json]
  -v, --verbose                   Be more verbose [default: false]
  -h, --help     Show this screen
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
	verbose := args["--verbose"].(bool)

	if verbose == true {
		log.Printf("Debug: args - %q", args)
	}

	var filename string
	if args["FILE"] == nil {
		filename = "-"
	} else {
		filename = args["FILE"].(string)
	}
	data, _ := readData(filename)

	var input_format string
	if args["--source"].(string) == "auto" {
		input_format = detectInputFormat(data)
	} else {
		input_format = args["--source"].(string)
	}
	parsed_data, _ := parseData(data, input_format)

	output_data, _ := outputData(parsed_data, args["--target"].(string))

	fmt.Println(string(output_data))
}
