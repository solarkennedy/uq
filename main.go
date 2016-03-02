package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)

	json, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(json))
}
