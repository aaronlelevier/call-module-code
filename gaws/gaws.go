package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// func callAwsApi() string {
// 	out, err := exec.Command("aws", "sts", "get-caller-identity").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return string(out)
// }

func getFilename() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	return path.Join(parent, "gaws/tests/data/ec2/describe-vpcs.json")
}

func readFile(filename string) []byte {
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return dat
	// fmt.Print(string(dat))
}

type Vpc struct {
	VpcId string
}

type Vpcs struct {
	Vpcs []Vpc
}

func decodeJson(b []byte) {
	var vpcs Vpcs
	json.Unmarshal(b, &vpcs)
	fmt.Printf("%+v\n", vpcs)
}

func main() {
	filename := getFilename()
	fmt.Printf("%s\n", filename)

	contents := readFile(filename)
	// fmt.Printf("%s\n", contents)

	decodeJson(contents)
}
