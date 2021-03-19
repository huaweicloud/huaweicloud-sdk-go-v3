package model

import (
	"encoding/json"

	"strings"
)

// 设备属性过滤信息，自定义结构。
type PropertyFilter struct {
	// 设备属性的路径信息，格式：service_id/DataProperty，例如门磁状态为“DoorWindow/status”。多个属性路径之间以逗号分隔。

	Path string `json:"path"`
	// 数据比较的操作符，当前支持的操作符有：>，<，>=，<=，=和between:表示数值区间，geo.circle.in:表示圆形区域范围内，geo.circle.out:表示圆形区域范围外。

	Operator string `json:"operator"`
	// 数据比较表达式的右值。与数据比较操作符between联用时，右值表示最小值和最大值，用逗号隔开，如“20,30”表示大于等于20小于30。

	Value string `json:"value"`

	Strategy *Strategy `json:"strategy,omitempty"`
}

func (o PropertyFilter) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PropertyFilter struct{}"
	}

	return strings.Join([]string{"PropertyFilter", string(data)}, " ")
}
