package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ColumnLineageV2 struct {

	// 上游血缘字段列表，列表大小：1至100
	InputColumns []ColumnDetails `json:"input_columns"`

	// 下游血缘字段列表，列表大小：1至100
	OutputColumns []ColumnDetails `json:"output_columns"`
}

func (o ColumnLineageV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnLineageV2 struct{}"
	}

	return strings.Join([]string{"ColumnLineageV2", string(data)}, " ")
}
