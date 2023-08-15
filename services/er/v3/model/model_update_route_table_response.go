package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteTableResponse Response Object
type UpdateRouteTableResponse struct {
	RouteTable *RouteTable `json:"route_table,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableResponse", string(data)}, " ")
}
