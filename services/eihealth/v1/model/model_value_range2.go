package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValueRange2 数值值域区间定义
type ValueRange2 struct {

	// 值域下限
	Lower *float32 `json:"lower,omitempty"`

	// 值域上限
	Upper *float32 `json:"upper,omitempty"`
}

func (o ValueRange2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueRange2 struct{}"
	}

	return strings.Join([]string{"ValueRange2", string(data)}, " ")
}
