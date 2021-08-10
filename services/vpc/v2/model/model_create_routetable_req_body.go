package model

import (
	"encoding/json"

	"strings"
)

//
type CreateRoutetableReqBody struct {
	Routetable *CreateRouteTableReq `json:"routetable"`
}

func (o CreateRoutetableReqBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"CreateRoutetableReqBody", string(data)}, " ")
}
