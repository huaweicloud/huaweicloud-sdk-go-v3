package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRouteTableResponse Response Object
type ShowRouteTableResponse struct {
	RouteTable *RouteTable `json:"route_table,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRouteTableResponse struct{}"
	}

	return strings.Join([]string{"ShowRouteTableResponse", string(data)}, " ")
}
