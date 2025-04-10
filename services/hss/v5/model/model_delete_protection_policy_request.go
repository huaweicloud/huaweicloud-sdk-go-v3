package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectionPolicyRequest Request Object
type DeleteProtectionPolicyRequest struct {

	// 租户企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 防护策略id
	PolicyId string `json:"policy_id"`
}

func (o DeleteProtectionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectionPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectionPolicyRequest", string(data)}, " ")
}
