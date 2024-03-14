package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleFieldDto struct {

	// 指标表对应字段名称。
	ColumnName *string `json:"column_name,omitempty"`

	// 指标表对应字段类型。
	ColumnType *string `json:"column_type,omitempty"`
}

func (o SimpleFieldDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleFieldDto struct{}"
	}

	return strings.Join([]string{"SimpleFieldDto", string(data)}, " ")
}
