package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CreateRoutetableReqBody struct {
	Routetable *CreateRouteTableReq `json:"routetable"`
}

func (o CreateRoutetableReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableReqBody struct{}"
	}

	return strings.Join([]string{"CreateRoutetableReqBody", string(data)}, " ")
}
