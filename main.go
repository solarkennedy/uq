package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {

	stat, _ := os.Stdin.Stat()
	var bytes []byte
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ = ioutil.ReadAll(os.Stdin)
	} else {
		//file, _ := os.Open(os.Args[len(os.Args)-1])
		bytes, _ = ioutil.ReadFile(os.Args[len(os.Args)-1])
	}

	json, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	} else {
		fmt.Println(string(json))
	}

}
