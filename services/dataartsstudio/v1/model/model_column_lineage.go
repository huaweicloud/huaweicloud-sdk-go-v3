package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ColumnLineage struct {

	// 作业算子名称
	Name string `json:"name"`

	// 上游血缘字段列表，列表大小：1至100
	InputColumns []ColumnDetails `json:"input_columns"`

	// 下游血缘字段列表，列表大小：1至100
	OutputColumns []ColumnDetails `json:"output_columns"`
}

func (o ColumnLineage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnLineage struct{}"
	}

	return strings.Join([]string{"ColumnLineage", string(data)}, " ")
}
