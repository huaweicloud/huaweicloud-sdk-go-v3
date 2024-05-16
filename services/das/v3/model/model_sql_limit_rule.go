package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlLimitRule SQL限流规则
type SqlLimitRule struct {

	// SQL限流规则ID
	Id string `json:"id"`

	// SQL类型
	SqlType string `json:"sql_type"`

	// 限流规则
	Pattern string `json:"pattern"`

	// 最大并发数
	MaxConcurrency int32 `json:"max_concurrency"`

	// 最大等待时间
	MaxWaiting *int32 `json:"max_waiting,omitempty"`
}

func (o SqlLimitRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlLimitRule struct{}"
	}

	return strings.Join([]string{"SqlLimitRule", string(data)}, " ")
}
