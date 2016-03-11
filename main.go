package main

import (
	"fmt"
	"io/ioutil"
	"os"

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
	}
	return 0
}

func main() {
	data := readData()
	exitcode := outputData(data)
	os.Exit(exitcode)
}
