package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseColumnDto 数据库列定义
type DatabaseColumnDto struct {

	// 列名
	Name string `json:"name"`

	Type *ColumnType `json:"type"`

	// 列描述信息
	Description *string `json:"description,omitempty"`

	// 列是否允许为空
	Nullable bool `json:"nullable"`

	// 是否作为主键
	Primary bool `json:"primary"`

	// 是否可查询
	Searchable bool `json:"searchable"`

	// 是否唯一
	Unique bool `json:"unique"`

	// 查询参数格式的提示信息
	Tips *string `json:"tips,omitempty"`
}

func (o DatabaseColumnDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseColumnDto struct{}"
	}

	return strings.Join([]string{"DatabaseColumnDto", string(data)}, " ")
}
