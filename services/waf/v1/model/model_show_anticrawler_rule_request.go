package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnticrawlerRuleRequest Request Object
type ShowAnticrawlerRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 规则id
	RuleId string `json:"rule_id"`
}

func (o ShowAnticrawlerRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnticrawlerRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAnticrawlerRuleRequest", string(data)}, " ")
}
