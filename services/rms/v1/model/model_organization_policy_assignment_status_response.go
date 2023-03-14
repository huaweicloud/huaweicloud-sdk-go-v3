package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组织合规规则部署状态。
type OrganizationPolicyAssignmentStatusResponse struct {

	// 组织合规规则ID。
	OrganizationPolicyAssignmentId *string `json:"organization_policy_assignment_id,omitempty"`

	// 组织合规规则名称。
	OrganizationPolicyAssignmentName *string `json:"organization_policy_assignment_name,omitempty"`

	// 组织合规规则部署状态。
	OrganizationPolicyAssignmentStatus *string `json:"organization_policy_assignment_status,omitempty"`

	// 当创建或更新组织合规规则失败时错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 当创建或更新组织合规规则失败时错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o OrganizationPolicyAssignmentStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyAssignmentStatusResponse struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyAssignmentStatusResponse", string(data)}, " ")
}
