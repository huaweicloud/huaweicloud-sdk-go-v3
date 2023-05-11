package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAggregateComplianceDetailsByPolicyAssignmentRequest struct {

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	Body *AggregateComplianceDetailRequest `json:"body,omitempty"`
}

func (o ShowAggregateComplianceDetailsByPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregateComplianceDetailsByPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"ShowAggregateComplianceDetailsByPolicyAssignmentRequest", string(data)}, " ")
}
