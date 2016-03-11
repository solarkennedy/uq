package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
)

func readData() []byte {
	var bytes []byte
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, _ = ioutil.ReadFile(os.Args[len(os.Args)-1])
	}
	return bytes
}

func outputData(data []byte) int {
	json, err := yaml.YAMLToJSON(data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return 1
	} else {
		fmt.Println(string(json))
		return 0
	}
}

func parseArgs() map[string]interface{} {
	usage := `uq - Universal serialized data reader to JSON

Usage:
  uq
  uq <file>
  uq --yaml
  uq -h | --help
  uq --version

Options:
  -h --help     Show this screen
  --version     Show version
  -y --yaml     Force reading input as YAML

Examples:
  cat input.yaml | uq | jq .
  jq --yaml input.yaml | jq .
`

	arguments, _ := docopt.Parse(usage, nil, true, "0.0.1", false)
	fmt.Println(arguments)
	return arguments
}

func main() {
	parseArgs()
	data := readData()
	exitcode := outputData(data)
	os.Exit(exitcode)
}
