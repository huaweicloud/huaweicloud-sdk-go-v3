package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Object2Map 外层 Map<String, Object>，键为字符串，值为包含 Metric2Map 和 string 的对象
type Object2Map struct {

	// 表示一个 Map<String, BigDecimal>，用于存储高精度数值类型的指标
	Metric map[string]float32 `json:"metric,omitempty"`

	// 目录状态相关描述
	Category *string `json:"category,omitempty"`
}

func (o Object2Map) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Object2Map struct{}"
	}

	return strings.Join([]string{"Object2Map", string(data)}, " ")
}
