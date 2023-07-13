package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteTableResponse Response Object
type CreateRouteTableResponse struct {
	RouteTable *RouteTable `json:"route_table,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTableResponse struct{}"
	}

	return strings.Join([]string{"CreateRouteTableResponse", string(data)}, " ")
}
