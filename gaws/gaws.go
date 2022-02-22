package main

import (
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

func main() {
	filename := getFilename()
	fmt.Printf("%s\n", filename)

	contents := readFile(filename)
	fmt.Printf("%s\n", contents)
}
