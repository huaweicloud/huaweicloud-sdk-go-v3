package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntelligentKillSessionHistory 限流信息
type IntelligentKillSessionHistory struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// kill开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// kill结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 操作历史下载链接
	DownloadLink *string `json:"download_link,omitempty"`
}

func (o IntelligentKillSessionHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntelligentKillSessionHistory struct{}"
	}

	return strings.Join([]string{"IntelligentKillSessionHistory", string(data)}, " ")
}
