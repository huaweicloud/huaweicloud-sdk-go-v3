package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisablePropagationRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	Body *PropagationRequestBody `json:"body,omitempty"`
}

func (o DisablePropagationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisablePropagationRequest struct{}"
	}

	return strings.Join([]string{"DisablePropagationRequest", string(data)}, " ")
}
