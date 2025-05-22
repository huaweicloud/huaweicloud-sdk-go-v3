package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RealTimeLogResponseBodyResult 结果
type RealTimeLogResponseBodyResult struct {

	// 是否还有剩余日志标识
	HasMoreData *bool `json:"has_more_data,omitempty"`

	// 偏移量，可用于下一次请求
	Offset *int32 `json:"offset,omitempty"`

	// 返回日志内容，可能会空，请再次请求
	Content *string `json:"content,omitempty"`

	// 当前请求偏移量
	CurrentOffset *int32 `json:"current_offset,omitempty"`
}

func (o RealTimeLogResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeLogResponseBodyResult struct{}"
	}

	return strings.Join([]string{"RealTimeLogResponseBodyResult", string(data)}, " ")
}
