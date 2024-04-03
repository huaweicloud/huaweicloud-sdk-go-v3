package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationPolicyAssignmentResponse Response Object
type UpdateOrganizationPolicyAssignmentResponse struct {

	// 组织合规规则创建者。
	OwnerId *string `json:"owner_id,omitempty"`

	// 组织ID。
	OrganizationId *string `json:"organization_id,omitempty"`

	// 组织合规规则资源唯一标识。
	OrganizationPolicyAssignmentUrn *string `json:"organization_policy_assignment_urn,omitempty"`

	// 组织合规规则ID。
	OrganizationPolicyAssignmentId *string `json:"organization_policy_assignment_id,omitempty"`

	// 组织合规规则名称。
	OrganizationPolicyAssignmentName *string `json:"organization_policy_assignment_name,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 触发周期。
	Period *string `json:"period,omitempty"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty"`

	// 规则参数。
	Parameters map[string]PolicyParameterValue `json:"parameters,omitempty"`

	// 策略ID。
	PolicyDefinitionId *string `json:"policy_definition_id,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateOrganizationPolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationPolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationPolicyAssignmentResponse", string(data)}, " ")
}
