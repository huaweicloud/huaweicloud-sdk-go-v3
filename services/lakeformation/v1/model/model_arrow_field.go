package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArrowField 定义了Arrow Field的结构，遵循Apache Arrow标准。
type ArrowField struct {

	// 字段名称。
	Name string `json:"name"`

	// 字段是否可为空。
	Nullable bool `json:"nullable"`

	// 字段类型。
	Type string `json:"type"`

	// 子字段列表（用于嵌套类型）。
	Children *[]ArrowField `json:"children,omitempty"`
}

func (o ArrowField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArrowField struct{}"
	}

	return strings.Join([]string{"ArrowField", string(data)}, " ")
}
