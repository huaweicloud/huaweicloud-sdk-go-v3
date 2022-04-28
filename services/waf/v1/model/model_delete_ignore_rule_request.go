package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteIgnoreRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	// 误报屏蔽规则id，您可以通过查询误报屏蔽规则列表（ListIgnoreRule）接口的响应体中的id字段获取误报屏蔽规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteIgnoreRuleRequest", string(data)}, " ")
}
