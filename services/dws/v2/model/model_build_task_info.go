package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildTaskInfo 用户配置build任务信息
type BuildTaskInfo struct {

	// 任务模式
	BuildMode string `json:"build_mode"`

	// 任务开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 策略ID
	CategoryId *string `json:"category_id,omitempty"`
}

func (o BuildTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildTaskInfo struct{}"
	}

	return strings.Join([]string{"BuildTaskInfo", string(data)}, " ")
}
