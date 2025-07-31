package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskStatisticsResponse Response Object
type ShowTaskStatisticsResponse struct {

	// 累计执行的任务总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 扫描中的任务数
	RunningNum *int32 `json:"running_num,omitempty"`

	// 最近一次扫描任务的创建时间
	LastTaskStartTime *int64 `json:"last_task_start_time,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowTaskStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskStatisticsResponse", string(data)}, " ")
}
