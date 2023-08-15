package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnticrawlerRuleRequest Request Object
type CreateAnticrawlerRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	Body *CreateAnticrawlerRuleRequestbody `json:"body,omitempty"`
}

func (o CreateAnticrawlerRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnticrawlerRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateAnticrawlerRuleRequest", string(data)}, " ")
}
