package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func readFile() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	parent := filepath.Dir(wd)
	return path.Join(parent, "gaws/tests/data/ec2/describe-vpcs.json")
}

func callAwsApi() string {
	out, err := exec.Command("aws", "sts", "get-caller-identity").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {
	out := readFile()
	fmt.Printf("%s\n", out)
}
