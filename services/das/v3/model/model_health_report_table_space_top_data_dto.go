package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceTopDataDto struct {

	// 库名
	DbName *string `json:"db_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 物理文件大小，单位MB
	File *float64 `json:"file,omitempty"`

	// 数据空间，单位MB
	Data *float64 `json:"data,omitempty"`

	// 碎片空间，单位MB
	Free *float64 `json:"free,omitempty"`

	// 碎片率
	FreeRate *float64 `json:"free_rate,omitempty"`

	// 索引空间，单位MB
	Index *float64 `json:"index,omitempty"`

	// 行数
	RowsCount *int64 `json:"rows_count,omitempty"`

	// 采集时间
	Time *int64 `json:"time,omitempty"`
}

func (o HealthReportTableSpaceTopDataDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceTopDataDto struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceTopDataDto", string(data)}, " ")
}
