package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTimeInfo 任务开始运行时间信息
type StartTimeInfo struct {

	// 任务启动时间
	StartTime string `json:"start_time"`

	// 任务运行频率（定时任务频率，每天运行还是单次运行）
	Frequency *string `json:"frequency,omitempty"`

	// 任务运行模式（按时间段运行还是按频率运行）
	Mode *string `json:"mode,omitempty"`

	// 每天运行时间段
	DayTimeFrame *[]TimeFrame `json:"day_time_frame,omitempty"`

	// 每次运行时间段
	SingleTimeFrame *[]TimeFrame `json:"single_time_frame,omitempty"`
}

func (o StartTimeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTimeInfo struct{}"
	}

	return strings.Join([]string{"StartTimeInfo", string(data)}, " ")
}
