package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组织合规规则请求体。
type OrganizationPolicyAssignmentRequest struct {

	// 需要排除配置规则的帐号。
	ExcludedAccounts *[]string `json:"excluded_accounts,omitempty"`

	// 组织合规规则名称。
	OrganizationPolicyAssignmentName string `json:"organization_policy_assignment_name"`

	ManagedPolicyAssignmentMetadata *ManagedPolicyAssignmentMetadata `json:"managed_policy_assignment_metadata,omitempty"`
}

func (o OrganizationPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyAssignmentRequest", string(data)}, " ")
}
