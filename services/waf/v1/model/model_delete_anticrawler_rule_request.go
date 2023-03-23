package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAnticrawlerRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteAnticrawlerRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnticrawlerRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAnticrawlerRuleRequest", string(data)}, " ")
}
