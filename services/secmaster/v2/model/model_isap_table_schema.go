package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapTableSchema 表架构
type IsapTableSchema struct {

	// 表格列列表
	Columns *[]IsapTableColumn `json:"columns,omitempty"`

	// 表主键列表
	PrimaryKey *[]string `json:"primary_key,omitempty"`

	// 表分区列表
	Partition *[]string `json:"partition,omitempty"`

	// 表水位列元素
	WatermarkColumn *string `json:"watermark_column,omitempty"`

	// 表水位延迟间隔
	WatermarkInterval *float32 `json:"watermark_interval,omitempty"`

	// 表时间过滤列
	TimeFilter *string `json:"time_filter,omitempty"`
}

func (o IsapTableSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableSchema struct{}"
	}

	return strings.Join([]string{"IsapTableSchema", string(data)}, " ")
}
