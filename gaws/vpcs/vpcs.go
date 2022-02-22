package vpcs

import (
	"encoding/json"
)

type Vpc struct {
	VpcId string
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
