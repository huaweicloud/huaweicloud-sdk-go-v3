package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtremumDto study作业最值信息
type ExtremumDto struct {

	// 最值
	Value *float64 `json:"value,omitempty"`

	// 最值所在的行数
	RowNumber *int64 `json:"row_number,omitempty"`

	// 最值所在的列名
	ColumnName *string `json:"column_name,omitempty"`

	// 最值所在的行名
	RowName *string `json:"row_name,omitempty"`
}

func (o ExtremumDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtremumDto struct{}"
	}

	return strings.Join([]string{"ExtremumDto", string(data)}, " ")
}
