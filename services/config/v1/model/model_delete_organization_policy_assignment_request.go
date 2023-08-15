package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationPolicyAssignmentRequest Request Object
type DeleteOrganizationPolicyAssignmentRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规规则ID。
	OrganizationPolicyAssignmentId string `json:"organization_policy_assignment_id"`
}

func (o DeleteOrganizationPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationPolicyAssignmentRequest", string(data)}, " ")
}
