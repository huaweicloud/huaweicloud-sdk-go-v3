package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipeSchema 管道模式
type PipeSchema struct {

	// 管道列
	Columns []TableColumnForIsapPipe `json:"columns"`

	// 管道水线列
	WatermarkColumn *string `json:"watermark_column,omitempty"`

	// 管道水线间隔时长
	WatermarkInterval *float32 `json:"watermark_interval,omitempty"`

	// 管道过滤列
	TimeFilter string `json:"time_filter"`
}

func (o PipeSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipeSchema struct{}"
	}

	return strings.Join([]string{"PipeSchema", string(data)}, " ")
}
