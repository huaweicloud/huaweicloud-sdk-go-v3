package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAntiTamperRuleRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id

	PolicyId string `json:"policy_id"`

	Body *CreateAntiTamperRulesRequestBody `json:"body,omitempty"`
}

func (o CreateAntiTamperRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiTamperRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateAntiTamperRuleRequest", string(data)}, " ")
}
