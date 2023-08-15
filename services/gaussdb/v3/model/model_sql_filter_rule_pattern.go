package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlFilterRulePattern SQL限流规则和最大并发数
type SqlFilterRulePattern struct {

	// SQL限流规则。
	Pattern string `json:"pattern"`

	// 最大并发数。
	MaxConcurrency int32 `json:"max_concurrency"`
}

func (o SqlFilterRulePattern) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlFilterRulePattern struct{}"
	}

	return strings.Join([]string{"SqlFilterRulePattern", string(data)}, " ")
}
