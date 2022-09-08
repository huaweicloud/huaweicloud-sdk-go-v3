package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateRouteTableRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	Body *AssociationRequestBody `json:"body,omitempty"`
}

func (o AssociateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"AssociateRouteTableRequest", string(data)}, " ")
}
