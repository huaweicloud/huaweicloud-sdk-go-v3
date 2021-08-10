package model

import (
	"encoding/json"

	"strings"
)

//
type AsscoiateReq struct {
	Subnets *AssociateRouteTableAndSubnetReq `json:"subnets"`
}

func (o AsscoiateReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AsscoiateReq struct{}"
	}

	return strings.Join([]string{"AsscoiateReq", string(data)}, " ")
}
