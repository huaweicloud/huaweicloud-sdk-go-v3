package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCcRuleRequest Request Object
type UpdateCcRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id（策略id从查询防护策略列表接口获取）
	PolicyId string `json:"policy_id"`

	// ID of the cc rule. It can be obtained by calling the **ListCcRules** API.
	RuleId string `json:"rule_id"`

	Body *UpdateCcRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCcRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateCcRuleRequest", string(data)}, " ")
}
