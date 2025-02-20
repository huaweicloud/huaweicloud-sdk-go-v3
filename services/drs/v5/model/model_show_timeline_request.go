package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTimelineRequest Request Object
type ShowTimelineRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage string `json:"X-Language"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTimelineRequest struct{}"
	}

	return strings.Join([]string{"ShowTimelineRequest", string(data)}, " ")
}
