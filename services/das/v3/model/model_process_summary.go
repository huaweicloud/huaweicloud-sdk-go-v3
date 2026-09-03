package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessSummary 会话概要信息
type ProcessSummary struct {

	// 概要参数名
	Key *string `json:"key,omitempty"`

	// 概要参数值
	Value *int64 `json:"value,omitempty"`
}

func (o ProcessSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessSummary struct{}"
	}

	return strings.Join([]string{"ProcessSummary", string(data)}, " ")
}
