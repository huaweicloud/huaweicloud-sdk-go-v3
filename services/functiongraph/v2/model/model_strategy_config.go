package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数策略配置。
type StrategyConfig struct {

	// 单函数最大实例数，v1取值0和-1，v2取值-1到1000
	Concurrency int32 `json:"concurrency"`

	// 单实例最大并发数，v2版本才支持，取值-1到200000
	ConcurrentNum int32 `json:"concurrent_num"`
}

func (o StrategyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StrategyConfig struct{}"
	}

	return strings.Join([]string{"StrategyConfig", string(data)}, " ")
}
