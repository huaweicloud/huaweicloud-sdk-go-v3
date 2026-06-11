package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArrowField Arrow Schema中的字段定义，包含字段名、类型、是否可空及元数据。
type ArrowField struct {

	// 字段名称。
	Name string `json:"name"`

	Type *ArrowType `json:"type"`

	// 字段是否允许为null。
	Nullable *bool `json:"nullable,omitempty"`

	// 字段的元数据信息。
	Metadata map[string]string `json:"metadata,omitempty"`

	Children *[]ArrowField `json:"children,omitempty"`
}

func (o ArrowField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArrowField struct{}"
	}

	return strings.Join([]string{"ArrowField", string(data)}, " ")
}
