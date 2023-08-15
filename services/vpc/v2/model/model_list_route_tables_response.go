package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRouteTablesResponse Response Object
type ListRouteTablesResponse struct {

	// 路由表对象列表
	Routetables    *[]RouteTableListResp `json:"routetables,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRouteTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRouteTablesResponse struct{}"
	}

	return strings.Join([]string{"ListRouteTablesResponse", string(data)}, " ")
}
