package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateStaticRouteRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	Body *CreateRouteRequestBody `json:"body,omitempty"`
}

func (o CreateStaticRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStaticRouteRequest struct{}"
	}

	return strings.Join([]string{"CreateStaticRouteRequest", string(data)}, " ")
}
