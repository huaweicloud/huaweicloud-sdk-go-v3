package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListRuleHitCountDto struct {

	// 规则id列表
	RuleIds []string `json:"rule_ids"`
}

func (o ListRuleHitCountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleHitCountDto struct{}"
	}

	return strings.Join([]string{"ListRuleHitCountDto", string(data)}, " ")
}
