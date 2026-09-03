package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndexUsageCondition 索引使用过滤条件
type IndexUsageCondition struct {

	// 过滤字段
	Field *string `json:"field,omitempty"`

	// 最小值
	MinValue *float64 `json:"min_value,omitempty"`

	// 最大值
	MaxValue *float64 `json:"max_value,omitempty"`
}

func (o IndexUsageCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndexUsageCondition struct{}"
	}

	return strings.Join([]string{"IndexUsageCondition", string(data)}, " ")
}
