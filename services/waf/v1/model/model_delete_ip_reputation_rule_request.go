package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpReputationRuleRequest Request Object
type DeleteIpReputationRuleRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 机房IP情报访问控制规则id，通过查询机房IP情报规则列表接口获取https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=WAF&api=ListIdcIpRule **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`
}

func (o DeleteIpReputationRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpReputationRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpReputationRuleRequest", string(data)}, " ")
}
