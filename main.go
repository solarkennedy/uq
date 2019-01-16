package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/clbanning/mxj"
	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
)

var version string

func readData(filename string) (bytes []byte, err error) {
	if filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}
	return bytes, err
}

func parseData(input_data []byte, format string) (parsed_data interface{}, err error) {
	if format == "json" {
		err = json.Unmarshal(input_data, &parsed_data)
	} else if format == "yaml" || format == "yml" {
		err = yaml.Unmarshal(input_data, &parsed_data)
	} else if format == "toml" {
		err = toml.Unmarshal(input_data, &parsed_data)
	} else if format == "xml" {
		// TODO: Is this the right way to handle xml?
		m := make(map[string]interface{})
		m, err = mxj.NewMapXml(input_data)
		return m, err
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
	} else if format == "toml" {
		var buf bytes.Buffer
		toml_enc := toml.NewEncoder(&buf)
		err = toml_enc.Encode(input_data)
		output_data = buf.Bytes()
	} else if format == "xml" {
		output_data, err = mxj.AnyXmlIndent(input_data, "", "\t")
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
  * toml
  * xml (Note: xml won't be a perfect conversion)
`

	arguments, _ := docopt.Parse(usage, nil, true, version, false)
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
	parsed_data, err := parseData(data, input_format)
	if err != nil {
		log.Fatal(err)
	}

	output_data, _ := outputData(parsed_data, args["--target"].(string))

	fmt.Println(string(output_data))
}
