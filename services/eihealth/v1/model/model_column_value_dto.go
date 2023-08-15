package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ColumnValueDto struct {

	// 列名
	Column string `json:"column"`

	// 该列对应的值
	Value *string `json:"value,omitempty"`
}

func (o ColumnValueDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnValueDto struct{}"
	}

	return strings.Join([]string{"ColumnValueDto", string(data)}, " ")
}
