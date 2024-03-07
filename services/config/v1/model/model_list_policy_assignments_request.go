package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyAssignmentsRequest Request Object
type ListPolicyAssignmentsRequest struct {

	// 合规规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListPolicyAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyAssignmentsRequest", string(data)}, " ")
}
