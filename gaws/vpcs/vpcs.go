package vpcs

import (
	"encoding/json"
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

// Decodes a lookup string and returns a map with these keys:
//
// - HName: handler name
// - FName: filter name
// - FValue: filter value
// - RKey: return key
func RegexMatch(t string) map[string]string {
	s := `(?P<HName>[a-z]+)\sName=(?P<FName>[A-Za-z\:]+),Values=(?P<FValue>[A-Za-z\-0-9]+)\s(?P<RKey>[A-Za-z]+)`
	r, _ := regexp.Compile(s)
	match := r.FindStringSubmatch(t)
	m := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = match[i]
		}
	}
	return m
}
