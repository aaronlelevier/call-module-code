package vpcs

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Tag struct {
	Key   string
	Value string
}

type Vpc struct {
	VpcId string
	Tags  []Tag
}

type Vpcs struct {
	Vpcs []Vpc
}

func DecodeVpcs(b []byte) Vpcs {
	var vpcs Vpcs
	json.Unmarshal(b, &vpcs)
	return vpcs
}

func ToVpc(v Vpcs) Vpc {
	return v.Vpcs[0]
}

func Tags(v Vpc) map[string]string {
	m := make(map[string]string)

	for _, t := range v.Tags {
		m[t.Key] = t.Value
	}
	return m
}

func IsMatch(v Vpc, k string, val string) bool {
	m := Tags(v)
	return m[k] == val
}

// awscli
// aws ec2 describe-vpcs --filters Name=tag:Default,Values=True

// should only find one 1 'vpc and return it's 'id'
// ${vpc Name=tag:Default,Values=True VpcId}

func RegexMatch() {
	// "p([a-z]+)ch"

	// s := "${vpc (.*?)}"
	// s := `([a-z]+)\sName=([A-Za-z/:]+),`
	s := `(?P<name>[a-z]+)\sName=(?P<fName>[A-Za-z\:]+),Values=(?P<fValues>[A-Za-z]+)\s(?P<rKey>[A-Za-z]+)`

	t := "${vpc Name=tag:Default,Values=True VpcId}"

	r, _ := regexp.Compile(s)
	fmt.Printf("is match: %+v\n", r.MatchString(t))
	fmt.Println(r.FindString(t))

	match := r.FindStringSubmatch(t)
	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	fmt.Printf("%s\n", result)
}
