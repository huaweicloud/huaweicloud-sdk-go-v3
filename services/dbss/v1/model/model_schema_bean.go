package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SchemaBean struct {

	// 目标字段
	Column *string `json:"column,omitempty"`

	// 目标数据库
	Schema *string `json:"schema,omitempty"`

	// 目标表
	Table *string `json:"table,omitempty"`
}

func (o SchemaBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaBean struct{}"
	}

	return strings.Join([]string{"SchemaBean", string(data)}, " ")
}
