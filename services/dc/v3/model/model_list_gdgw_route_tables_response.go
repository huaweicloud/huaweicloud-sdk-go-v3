package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGdgwRouteTablesResponse Response Object
type ListGdgwRouteTablesResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 全域接入网关路由表
	GdgwRoutetable *[]ShowGdgwRoutetable `json:"gdgw_routetable,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListGdgwRouteTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGdgwRouteTablesResponse struct{}"
	}

	return strings.Join([]string{"ListGdgwRouteTablesResponse", string(data)}, " ")
}
