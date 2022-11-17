package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleHitCountObject struct {

	// 规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 规则击中次数
	RuleHitCount *int32 `json:"rule_hit_count,omitempty"`
}

func (o RuleHitCountObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleHitCountObject struct{}"
	}

	return strings.Join([]string{"RuleHitCountObject", string(data)}, " ")
}
