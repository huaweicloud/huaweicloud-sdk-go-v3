package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationPolicyAssignmentRequest Request Object
type UpdateOrganizationPolicyAssignmentRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规规则ID。
	OrganizationPolicyAssignmentId string `json:"organization_policy_assignment_id"`

	Body *OrganizationPolicyAssignmentRequest `json:"body,omitempty"`
}

func (o UpdateOrganizationPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationPolicyAssignmentRequest", string(data)}, " ")
}
