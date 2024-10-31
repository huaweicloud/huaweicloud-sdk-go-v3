package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearAccessLogRuleHitCountsDto clear access log rule hit counts dto
type ClearAccessLogRuleHitCountsDto struct {

	// 删除规则击中次数请求的规则列表，规则id可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。
	RuleIds []string `json:"rule_ids"`
}

func (o ClearAccessLogRuleHitCountsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearAccessLogRuleHitCountsDto struct{}"
	}

	return strings.Join([]string{"ClearAccessLogRuleHitCountsDto", string(data)}, " ")
}
