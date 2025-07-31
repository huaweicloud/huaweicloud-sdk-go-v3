package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListHandleAffectBaselineRequestBodyCheckRuleList struct {

	// **参数解释** 基线检查的名称 **约束限制** 不涉及 **取值范围**   字符长度0-256位 **默认取值** 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释** 检查项id **约束限制** 不涉及 **取值范围**   字符长度0-256位 **默认取值** 不涉及
	CheckRuleId *string `json:"check_rule_id,omitempty"`

	// **参数解释** 基线检查标准类型 **约束限制** 不涉及 **取值范围**   - cn_standard : 等保合规标准 - hw_standard : 云安全实践标准 - cis_standard : 通用安全标准 **默认取值** 不涉及
	Standard *string `json:"standard,omitempty"`
}

func (o ListHandleAffectBaselineRequestBodyCheckRuleList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHandleAffectBaselineRequestBodyCheckRuleList struct{}"
	}

	return strings.Join([]string{"ListHandleAffectBaselineRequestBodyCheckRuleList", string(data)}, " ")
}
