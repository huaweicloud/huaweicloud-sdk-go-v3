package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCcRuleRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id（策略id从查询防护策略列表接口获取）
	PolicyId string `json:"policy_id"`

	Body *CreateCcRuleRequestBody `json:"body,omitempty"`
}

func (o CreateCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCcRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateCcRuleRequest", string(data)}, " ")
}
