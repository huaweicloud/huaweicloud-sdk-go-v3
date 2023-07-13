package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRouteTableRequest Request Object
type DeleteRouteTableRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`
}

func (o DeleteRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteRouteTableRequest", string(data)}, " ")
}
