package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePolicyProtectHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// 域名id9从获取防护网站列表获取域名id）

	Hosts string `json:"hosts"`
}

func (o UpdatePolicyProtectHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyProtectHostRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyProtectHostRequest", string(data)}, " ")
}
