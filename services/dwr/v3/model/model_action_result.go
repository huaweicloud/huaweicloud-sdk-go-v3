package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点错误处理的定义
type ActionResult struct {

	// 触发错误处理需符合的条件
	Match string `json:"match"`

	// 每次重试间隔时间
	RetryInterval int32 `json:"retry_interval"`

	// 最多重试次数
	MaxRetry *int32 `json:"max_retry,omitempty"`

	// 下一个状态
	NextState string `json:"next_state"`

	// 是否为结束状态
	IsTerminal *bool `json:"is_terminal,omitempty"`
}

func (o ActionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionResult struct{}"
	}

	return strings.Join([]string{"ActionResult", string(data)}, " ")
}
