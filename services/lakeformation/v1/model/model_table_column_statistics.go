package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableColumnStatistics 表中列的统计
type TableColumnStatistics struct {
	ColumnStatisticsDesc *TableColumnStatisticsDescription `json:"column_statistics_desc"`

	// 列统计信息
	ColumnStatisticsObjects []ColumnStatisticsObj `json:"column_statistics_objects"`
}

func (o TableColumnStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableColumnStatistics struct{}"
	}

	return strings.Join([]string{"TableColumnStatistics", string(data)}, " ")
}
