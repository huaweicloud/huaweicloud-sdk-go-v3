package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomizeParameter 自定义参数
type CustomizeParameter struct {

	// 别名。
	Alias string `json:"alias"`

	// 行为类型。
	BehaviorType string `json:"behavior_type"`

	// 阈值。
	Threshold *float64 `json:"threshold,omitempty"`

	// 去重。
	Deduplication string `json:"deduplication"`
}

func (o CustomizeParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeParameter struct{}"
	}

	return strings.Join([]string{"CustomizeParameter", string(data)}, " ")
}
