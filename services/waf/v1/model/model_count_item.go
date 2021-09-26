package model

import (
	"encoding/json"

	"strings"
)

// 攻击事件统计结果
type CountItem struct {
	// 类型

	Key *string `json:"key,omitempty"`
	// 数量

	Num *int32 `json:"num,omitempty"`
}

func (o CountItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CountItem struct{}"
	}

	return strings.Join([]string{"CountItem", string(data)}, " ")
}
