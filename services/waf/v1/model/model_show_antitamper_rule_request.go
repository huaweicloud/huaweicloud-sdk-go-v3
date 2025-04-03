package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntitamperRuleRequest Request Object
type ShowAntitamperRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 防篡改规则id，通过查询防篡改规则列表接口（ListAntitamperRule）获取
	RuleId string `json:"rule_id"`
}

func (o ShowAntitamperRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntitamperRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAntitamperRuleRequest", string(data)}, " ")
}
