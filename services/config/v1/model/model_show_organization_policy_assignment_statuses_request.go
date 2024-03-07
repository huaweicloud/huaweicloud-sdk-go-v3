package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyAssignmentStatusesRequest Request Object
type ShowOrganizationPolicyAssignmentStatusesRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规规则ID
	OrganizationPolicyAssignmentId *string `json:"organization_policy_assignment_id,omitempty"`

	// 组织合规规则名称。
	OrganizationPolicyAssignmentName *string `json:"organization_policy_assignment_name,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ShowOrganizationPolicyAssignmentStatusesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyAssignmentStatusesRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyAssignmentStatusesRequest", string(data)}, " ")
}
