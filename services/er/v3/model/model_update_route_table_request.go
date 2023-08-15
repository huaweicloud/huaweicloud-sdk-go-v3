package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteTableRequest Request Object
type UpdateRouteTableRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	Body *UpdateRouteTableRequestBody `json:"body,omitempty"`
}

func (o UpdateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableRequest", string(data)}, " ")
}
