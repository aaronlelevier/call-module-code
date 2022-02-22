package main

import (
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

func main() {

	// test payload name
	name := os.Args[1]

	filename := getFilename(name)
	fmt.Printf("%s\n", filename)

	contents := readFile(filename)
	// fmt.Printf("%s\n", contents)

	vpcs0 := vpcs.DecodeVpcs(contents)
	fmt.Printf("%+v\n", vpcs0)

	vpc := vpcs.ToVpc(vpcs0)
	fmt.Printf("%+v\n", vpc.VpcId)
}
