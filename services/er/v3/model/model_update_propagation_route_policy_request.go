package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePropagationRoutePolicyRequest Request Object
type UpdatePropagationRoutePolicyRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 传播ID
	PropagationId string `json:"propagation_id"`

	Body *UpdatePropagationRequestBody `json:"body,omitempty"`
}

func (o UpdatePropagationRoutePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePropagationRoutePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePropagationRoutePolicyRequest", string(data)}, " ")
}
