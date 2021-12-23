package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePolicyProtectHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id

	PolicyId string `json:"policy_id"`
	// 域名id，您可以通过调用查询云模式防护域名列表（ListHost）获取域名id

	Hosts string `json:"hosts"`
}

func (o UpdatePolicyProtectHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyProtectHostRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyProtectHostRequest", string(data)}, " ")
}
