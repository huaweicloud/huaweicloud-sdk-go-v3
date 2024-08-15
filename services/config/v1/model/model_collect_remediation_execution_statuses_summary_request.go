package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectRemediationExecutionStatusesSummaryRequest Request Object
type CollectRemediationExecutionStatusesSummaryRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	Body *RemediationExecutionStatusesSummaryRequestBody `json:"body,omitempty"`
}

func (o CollectRemediationExecutionStatusesSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectRemediationExecutionStatusesSummaryRequest struct{}"
	}

	return strings.Join([]string{"CollectRemediationExecutionStatusesSummaryRequest", string(data)}, " ")
}
