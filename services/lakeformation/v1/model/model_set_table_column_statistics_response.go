package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTableColumnStatisticsResponse Response Object
type SetTableColumnStatisticsResponse struct {
	ColumnStatisticsDesc *TableColumnStatisticsDescription `json:"column_statistics_desc,omitempty"`

	// 列统计信息
	ColumnStatisticsObjects *[]ColumnStatisticsObj `json:"column_statistics_objects,omitempty"`
	HttpStatusCode          int                    `json:"-"`
}

func (o SetTableColumnStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTableColumnStatisticsResponse struct{}"
	}

	return strings.Join([]string{"SetTableColumnStatisticsResponse", string(data)}, " ")
}
