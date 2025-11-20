package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpReputationRuleRequest Request Object
type UpdateIpReputationRuleRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符  **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id，响应体的id字段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 机房IP情报访问控制规则id，通过\"查询威胁情报规则列表\"接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`

	Body *UpdateIpReputationRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateIpReputationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpReputationRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpReputationRuleRequest", string(data)}, " ")
}
