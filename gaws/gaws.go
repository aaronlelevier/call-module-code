package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"example.com/gaws/vpcs"
)

// func callAwsApi() string {
// 	out, err := exec.Command("aws", "sts", "get-caller-identity").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return string(out)
// }

func getFilename(part string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	return path.Join(parent, "gaws/tests/data/ec2", part)
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

func decodeVpcs(b []byte) Vpcs {
	var vpcs Vpcs
	json.Unmarshal(b, &vpcs)
	return vpcs
}

func main() {

	// test payload name
	name := os.Args[1]

	vpcs.Hello()

	filename := getFilename(name)
	fmt.Printf("%s\n", filename)

	contents := readFile(filename)
	// fmt.Printf("%s\n", contents)

	vpcs := decodeVpcs(contents)
	fmt.Printf("%+v\n", vpcs)

	vpcId := vpcs.Vpcs[0].VpcId
	fmt.Printf("%+v\n", vpcId)
}
