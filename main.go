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
	fmt.Println("args:")
	fmt.Println(args)
	fmt.Println("Now data:")
	data := readData()
	exitcode := outputData(data)
	os.Exit(exitcode)
}
