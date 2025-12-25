package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Object1Map 外层 Map<String, Object>，键为字符串，值为包含 Metric1Map 和 string 的对象
type Object1Map struct {

	// 表示一个 Map<String, Integer>，用于存储整数类型的指标
	Metric map[string]int32 `json:"metric,omitempty"`

	// 目录状态相关描述
	Category *string `json:"category,omitempty"`
}

func (o Object1Map) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Object1Map struct{}"
	}

	return strings.Join([]string{"Object1Map", string(data)}, " ")
}
