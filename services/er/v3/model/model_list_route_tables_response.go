package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRouteTablesResponse struct {

	// 路由表列表
	RouteTables *[]RouteTable `json:"route_tables,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRouteTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRouteTablesResponse struct{}"
	}

	return strings.Join([]string{"ListRouteTablesResponse", string(data)}, " ")
}
