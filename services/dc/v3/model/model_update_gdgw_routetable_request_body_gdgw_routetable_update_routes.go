package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableUpdateRoutes struct {

	// 目的地址
	Destination string `json:"destination"`

	// 下一跳
	Nexthop string `json:"nexthop"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateGdgwRoutetableRequestBodyGdgwRoutetableUpdateRoutes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRoutetableRequestBodyGdgwRoutetableUpdateRoutes struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRoutetableRequestBodyGdgwRoutetableUpdateRoutes", string(data)}, " ")
}
