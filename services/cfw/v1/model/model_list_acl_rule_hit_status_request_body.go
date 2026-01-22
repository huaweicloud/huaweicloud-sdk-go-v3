package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAclRuleHitStatusRequestBody struct {

	// **参数解释**： 规则ID，可通过查询防护规则接口查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RuleIds []string `json:"rule_ids"`
}

func (o ListAclRuleHitStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclRuleHitStatusRequestBody struct{}"
	}

	return strings.Join([]string{"ListAclRuleHitStatusRequestBody", string(data)}, " ")
}
