package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationPolicyRequest Request Object
type DeleteOrganizationPolicyRequest struct {

	// 资源策略ID
	OrganizationPolicyId string `json:"organization_policy_id"`
}

func (o DeleteOrganizationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationPolicyRequest", string(data)}, " ")
}
