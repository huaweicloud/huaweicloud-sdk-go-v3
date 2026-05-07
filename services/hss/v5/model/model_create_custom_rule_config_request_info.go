package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCustomRuleConfigRequestInfo struct {

	// **参数解释**： 规则名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
	RuleName string `json:"rule_name"`

	// **参数解释**: 是否选择所有主机 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否  **默认取值**: false
	IsAllHost *bool `json:"is_all_host,omitempty"`

	CustomRuleValueInfo *CustomRuleValueInfo `json:"custom_rule_value_info"`

	// **参数解释**: agent列表 **约束限制**: 不涉及 **取值范围**: 1-1000个agentID **默认取值**: 不涉及
	AgentIds *[]string `json:"agent_ids,omitempty"`
}

func (o CreateCustomRuleConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRuleConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateCustomRuleConfigRequestInfo", string(data)}, " ")
}
