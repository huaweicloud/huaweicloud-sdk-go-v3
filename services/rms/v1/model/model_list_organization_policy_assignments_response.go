package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPolicyAssignmentsResponse Response Object
type ListOrganizationPolicyAssignmentsResponse struct {

	// 组织合规规则列表。
	OrganizationPolicyAssignments *[]OrganizationPolicyAssignmentResponse `json:"organization_policy_assignments,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOrganizationPolicyAssignmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPolicyAssignmentsResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationPolicyAssignmentsResponse", string(data)}, " ")
}
