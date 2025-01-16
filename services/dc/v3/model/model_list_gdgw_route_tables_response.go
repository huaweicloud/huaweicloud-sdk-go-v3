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
	GdgwRoutetables *[]CommonRoutetable `json:"gdgw_routetables,omitempty"`

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGdgwRouteTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGdgwRouteTablesResponse struct{}"
	}

	return strings.Join([]string{"ListGdgwRouteTablesResponse", string(data)}, " ")
}
