package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationPolicyRequest Request Object
type UpdateOrganizationPolicyRequest struct {

	// 组织策略ID
	OrganizationPolicyId string `json:"organization_policy_id"`

	Body *OrganizationPolicyUpdateReq `json:"body,omitempty"`
}

func (o UpdateOrganizationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationPolicyRequest", string(data)}, " ")
}
