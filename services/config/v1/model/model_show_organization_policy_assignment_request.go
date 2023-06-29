package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyAssignmentRequest Request Object
type ShowOrganizationPolicyAssignmentRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	// 组织合规规则ID。
	OrganizationPolicyAssignmentId string `json:"organization_policy_assignment_id"`
}

func (o ShowOrganizationPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyAssignmentRequest", string(data)}, " ")
}
