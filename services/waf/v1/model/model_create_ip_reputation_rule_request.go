package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpReputationRuleRequest Request Object
type CreateIpReputationRuleRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	Body *CreateIpReputationRuleRequestBody `json:"body,omitempty"`
}

func (o CreateIpReputationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpReputationRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateIpReputationRuleRequest", string(data)}, " ")
}
