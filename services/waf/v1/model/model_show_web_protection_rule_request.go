package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebProtectionRuleRequest Request Object
type ShowWebProtectionRuleRequest struct {

	// **参数解释：** 语言 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释：** 防护规则ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`
}

func (o ShowWebProtectionRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebProtectionRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowWebProtectionRuleRequest", string(data)}, " ")
}
