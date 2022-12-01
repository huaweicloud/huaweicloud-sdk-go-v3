package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务开始运行时间信息
type StartTimeInfo struct {

	// 任务启动时间
	StartTime string `json:"start_time"`

	Frequency *string `json:"frequency,omitempty"`

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
