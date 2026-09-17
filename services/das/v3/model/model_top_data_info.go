package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopDataInfo TopDataInfo对象
type TopDataInfo struct {

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
	RowsCount *float64 `json:"rows_count,omitempty"`

	// 采集时间（Unix timestamp），单位：毫秒
	Time *float64 `json:"time,omitempty"`

	// 增长量，单位MB
	Growth *float64 `json:"growth,omitempty"`
}

func (o TopDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopDataInfo struct{}"
	}

	return strings.Join([]string{"TopDataInfo", string(data)}, " ")
}
