package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数策略配置。
type StrategyConfig struct {
	// 0：函数被禁用;-1：函数被启用。

	Concurrency int32 `json:"concurrency"`
	// 函数并发数

	ConcurrentNum *int32 `json:"concurrent_num,omitempty"`
}

func (o StrategyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StrategyConfig struct{}"
	}

	return strings.Join([]string{"StrategyConfig", string(data)}, " ")
}
