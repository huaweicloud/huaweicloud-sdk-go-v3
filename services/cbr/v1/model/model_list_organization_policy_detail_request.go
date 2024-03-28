package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPolicyDetailRequest Request Object
type ListOrganizationPolicyDetailRequest struct {

	// 组织策略ID
	OrganizationPolicyId string `json:"organization_policy_id"`
}

func (o ListOrganizationPolicyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPolicyDetailRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationPolicyDetailRequest", string(data)}, " ")
}
