package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePipeSchemaResponse Response Object
type UpdatePipeSchemaResponse struct {

	// 管道列
	Columns *[]TableColumnForIsapPipe `json:"columns,omitempty"`

	// 管道水线列
	WatermarkColumn *string `json:"watermark_column,omitempty"`

	// 管道水线间隔时长
	WatermarkInterval *float32 `json:"watermark_interval,omitempty"`

	// 管道过滤列
	TimeFilter     *string `json:"time_filter,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePipeSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePipeSchemaResponse struct{}"
	}

	return strings.Join([]string{"UpdatePipeSchemaResponse", string(data)}, " ")
}
