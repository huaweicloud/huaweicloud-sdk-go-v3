package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPolicyAssignmentsRequest Request Object
type ListOrganizationPolicyAssignmentsRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规规则名称。
	OrganizationPolicyAssignmentName *string `json:"organization_policy_assignment_name,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListOrganizationPolicyAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPolicyAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationPolicyAssignmentsRequest", string(data)}, " ")
}
