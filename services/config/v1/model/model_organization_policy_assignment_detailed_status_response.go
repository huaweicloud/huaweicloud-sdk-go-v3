package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationPolicyAssignmentDetailedStatusResponse 组织合规规则部署详细状态。
type OrganizationPolicyAssignmentDetailedStatusResponse struct {

	// 帐号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 合规规则ID。
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty"`

	// 合规规则名称。
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 成员帐号中配置规则的部署状态。
	MemberAccountPolicyAssignmentStatus *string `json:"member_account_policy_assignment_status,omitempty"`

	// 当创建或更新合规规则失败时错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 当创建或更新合规规则失败时错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o OrganizationPolicyAssignmentDetailedStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyAssignmentDetailedStatusResponse struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyAssignmentDetailedStatusResponse", string(data)}, " ")
}
