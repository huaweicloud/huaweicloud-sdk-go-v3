package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRuleHitCountDto struct {

	// 规则id列表，规则ID，可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。
	RuleIds []string `json:"rule_ids"`
}

func (o ListRuleHitCountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleHitCountDto struct{}"
	}

	return strings.Join([]string{"ListRuleHitCountDto", string(data)}, " ")
}
