package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateRoutetableReqBody struct {
	Routetable *UpdateRouteTableReq `json:"routetable"`
}

func (o UpdateRoutetableReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableReqBody", string(data)}, " ")
}
