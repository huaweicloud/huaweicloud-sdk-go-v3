package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetTableColumnStatisticsInput 获取表列统计信息输入参数
type GetTableColumnStatisticsInput struct {

	// 列名
	ColumnNames []string `json:"column_names"`
}

func (o GetTableColumnStatisticsInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTableColumnStatisticsInput struct{}"
	}

	return strings.Join([]string{"GetTableColumnStatisticsInput", string(data)}, " ")
}
