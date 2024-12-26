package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGdgwRouteTableResponse Response Object
type UpdateGdgwRouteTableResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 全域接入网关路由表
	GdgwRoutetable *[]CommonRoutetable `json:"gdgw_routetable,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateGdgwRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRouteTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRouteTableResponse", string(data)}, " ")
}
