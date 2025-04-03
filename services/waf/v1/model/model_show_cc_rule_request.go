package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCcRuleRequest Request Object
type ShowCcRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id（策略id从查询防护策略列表接口获取）
	PolicyId string `json:"policy_id"`

	// ID of the cc rule. It can be obtained by calling the **ListCcRules** API.
	RuleId string `json:"rule_id"`
}

func (o ShowCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCcRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowCcRuleRequest", string(data)}, " ")
}
