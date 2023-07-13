package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomRuleRequest Request Object
type CreateCustomRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护策略id，通过指定防护策略id来指明查询该防护策略下的防护规则，您可以通过调用查询防护策略列表（ListPolicy）获取策略id
	PolicyId string `json:"policy_id"`

	Body *CreateCustomRuleRequestBody `json:"body,omitempty"`
}

func (o CreateCustomRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomRuleRequest", string(data)}, " ")
}
