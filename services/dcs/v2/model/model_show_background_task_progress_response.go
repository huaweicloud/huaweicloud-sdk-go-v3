package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackgroundTaskProgressResponse Response Object
type ShowBackgroundTaskProgressResponse struct {

	// 总体进度,百分比
	Progress *int32 `json:"progress,omitempty"`

	// 剩余时间(单位秒)
	RemainTime *int32 `json:"remain_time,omitempty"`

	// 任务详情列表
	StepDetails    *[]StepDetail `json:"step_details,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowBackgroundTaskProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundTaskProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowBackgroundTaskProgressResponse", string(data)}, " ")
}
