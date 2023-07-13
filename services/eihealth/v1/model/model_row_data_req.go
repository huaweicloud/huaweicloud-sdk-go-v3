package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RowDataReq struct {

	// 列及对应值列表
	ColumnValues []ColumnValueDto `json:"column_values"`
}

func (o RowDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RowDataReq struct{}"
	}

	return strings.Join([]string{"RowDataReq", string(data)}, " ")
}
