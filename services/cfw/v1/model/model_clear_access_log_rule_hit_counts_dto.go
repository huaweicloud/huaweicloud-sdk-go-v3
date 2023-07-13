package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearAccessLogRuleHitCountsDto clear access log rule hit counts dto
type ClearAccessLogRuleHitCountsDto struct {

	// 规则id列表
	RuleIds []string `json:"rule_ids"`
}

func (o ClearAccessLogRuleHitCountsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearAccessLogRuleHitCountsDto struct{}"
	}

	return strings.Join([]string{"ClearAccessLogRuleHitCountsDto", string(data)}, " ")
}
