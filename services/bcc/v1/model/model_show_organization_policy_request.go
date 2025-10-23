package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyRequest Request Object
type ShowOrganizationPolicyRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 组织策略Id
	OrganizationPolicyId string `json:"organization_policy_id"`
}

func (o ShowOrganizationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyRequest", string(data)}, " ")
}
