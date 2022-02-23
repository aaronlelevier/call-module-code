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
	tagKey := os.Args[2]
	tagValue := os.Args[3]

	filename := getFilename(name)
	fmt.Printf("%s\n", filename)
	fmt.Printf("%s\n", tagKey)
	fmt.Printf("%s\n", tagValue)

	contents := readFile(filename)
	// fmt.Printf("%s\n", contents)

	vpcs0 := vpcs.DecodeVpcs(contents)
	fmt.Printf("%+v\n", vpcs0)

	vpc := vpcs.ToVpc(vpcs0)
	fmt.Printf("%+v\n", vpc.VpcId)

	tags := vpcs.Tags(vpc)
	fmt.Printf("%+v\n", tags)

	isMatch := vpcs.IsMatch(vpc, tagKey, tagValue)
	fmt.Printf("isMatch: %+v\n", isMatch)

	t := "${vpc Name=tag:Default,Values=True VpcId}"
	m := vpcs.RegexMatch(t)
	fmt.Printf("%s\n", m)
}
