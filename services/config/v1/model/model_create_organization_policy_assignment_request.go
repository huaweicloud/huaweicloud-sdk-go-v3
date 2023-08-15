package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationPolicyAssignmentRequest Request Object
type CreateOrganizationPolicyAssignmentRequest struct {

	// 组织ID。
	OrganizationId string `json:"organization_id"`

	Body *OrganizationPolicyAssignmentRequest `json:"body,omitempty"`
}

func (o CreateOrganizationPolicyAssignmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"CreateOrganizationPolicyAssignmentRequest", string(data)}, " ")
}
