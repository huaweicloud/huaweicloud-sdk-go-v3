package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 错误处理策略详情
type OnError struct {

	// 错误匹配表达式，用来过滤需要处理的异常
	Error *string `json:"error,omitempty"`

	// 下一步骤节点ID
	Transition *string `json:"transition,omitempty"`

	// 重试策略名称
	RetryRef *string `json:"retry_ref,omitempty"`
}

func (o OnError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OnError struct{}"
	}

	return strings.Join([]string{"OnError", string(data)}, " ")
}
