package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAppLogResult 部署记录执行日志结果
type QueryAppLogResult struct {

	// 是否有更多日志信息
	HasMore *bool `json:"has_more,omitempty"`

	// 日志信息
	Text *string `json:"text,omitempty"`

	// 日志起始偏移量
	Offset *string `json:"offset,omitempty"`

	// 日志末位偏移量，与offset参数配合使用
	EndOffset *string `json:"end_offset,omitempty"`
}

func (o QueryAppLogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAppLogResult struct{}"
	}

	return strings.Join([]string{"QueryAppLogResult", string(data)}, " ")
}
