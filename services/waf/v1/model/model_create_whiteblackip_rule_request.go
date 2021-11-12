package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWhiteblackipRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`

	Body *CreateWhiteBlackIpRuleRequestBody `json:"body,omitempty"`
}

func (o CreateWhiteblackipRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhiteblackipRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateWhiteblackipRuleRequest", string(data)}, " ")
}
