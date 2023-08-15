package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoutetableReqBody
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
