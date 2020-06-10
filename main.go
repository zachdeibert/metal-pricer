package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/zachdeibert/metal-pricer/metalbythefoot"
)

func main() {
	res, err := metalbythefoot.ScanPrices()
	if err != nil {
		panic(err)
	}
	bin, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile("db.json", bin, 0755); err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile("db.js", []byte(fmt.Sprintf("r(%s);\n", string(bin))), 0755); err != nil {
		panic(err)
	}
	fmt.Println("File written.")
}
