package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareSlowLogTemplatesRequestBody 对比慢日志模板列表请求体
type CompareSlowLogTemplatesRequestBody struct {

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`

	// 对比日期开始时间（Unix timestamp），单位：毫秒
	ComparativeStartTime *int64 `json:"comparative_start_time,omitempty"`

	// 对比日期结束时间（Unix timestamp），单位：毫秒
	ComparativeEndTime *int64 `json:"comparative_end_time,omitempty"`

	// 基线日期开始时间（Unix timestamp），单位：毫秒
	BaseLineStartTime *int64 `json:"base_line_start_time,omitempty"`

	// 基线日期结束时间（Unix timestamp），单位：毫秒
	BaseLineEndTime *int64 `json:"base_line_end_time,omitempty"`
}

func (o CompareSlowLogTemplatesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareSlowLogTemplatesRequestBody struct{}"
	}

	return strings.Join([]string{"CompareSlowLogTemplatesRequestBody", string(data)}, " ")
}
