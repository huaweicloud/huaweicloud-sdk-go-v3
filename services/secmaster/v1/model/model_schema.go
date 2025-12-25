package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Schema 格式
type Schema struct {

	// 列
	Columns *[]SchemaColumns `json:"columns,omitempty"`

	// 分区
	Partition *[]string `json:"partition,omitempty"`

	// 主键
	PrimaryKey *[]string `json:"primary_key,omitempty"`

	// 时间过滤列
	TimeFilter *string `json:"time_filter,omitempty"`

	// 水印列
	WatermarkColumn *string `json:"watermark_column,omitempty"`

	// 水印间隔
	WatermarkInterval float32 `json:"watermark_interval,omitempty"`
}

func (o Schema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Schema struct{}"
	}

	return strings.Join([]string{"Schema", string(data)}, " ")
}
