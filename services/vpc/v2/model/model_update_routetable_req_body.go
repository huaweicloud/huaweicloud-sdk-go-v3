package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateRoutetableReqBody struct {
	Routetable *UpdateRouteTableReq `json:"routetable"`
}

func (o UpdateRoutetableReqBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableReqBody", string(data)}, " ")
}
