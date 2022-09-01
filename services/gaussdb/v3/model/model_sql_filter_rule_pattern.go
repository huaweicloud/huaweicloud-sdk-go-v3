package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SQL限流规则和最大并发数
type SqlFilterRulePattern struct {

	// SQL限流规则。
	Pattern string `json:"pattern" xml:"pattern"`

	// 最大并发数。
	MaxConcurrency int32 `json:"max_concurrency" xml:"max_concurrency"`
}

func (o SqlFilterRulePattern) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlFilterRulePattern struct{}"
	}

	return strings.Join([]string{"SqlFilterRulePattern", string(data)}, " ")
}
