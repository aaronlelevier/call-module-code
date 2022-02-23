package vpcs

import (
	"encoding/json"
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
