package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearAccessLogRuleHitCountsDto clear access log rule hit counts dto
type ClearAccessLogRuleHitCountsDto struct {

	// **参数解释**： 删除规则击中次数请求的规则列表，规则ID可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	RuleIds []string `json:"rule_ids"`
}

func (o ClearAccessLogRuleHitCountsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearAccessLogRuleHitCountsDto struct{}"
	}

	return strings.Join([]string{"ClearAccessLogRuleHitCountsDto", string(data)}, " ")
}
