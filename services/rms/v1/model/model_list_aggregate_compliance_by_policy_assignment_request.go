package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAggregateComplianceByPolicyAssignmentRequest struct {

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	Body *AggregatePolicyAssignmentsRequest `json:"body,omitempty"`
}

func (o ListAggregateComplianceByPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregateComplianceByPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"ListAggregateComplianceByPolicyAssignmentRequest", string(data)}, " ")
}
