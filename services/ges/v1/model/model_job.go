package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务
type Job struct {
	// 任务ID。

	JobId string `json:"jobId"`
	// 任务状态。  - pending:等待中 - running:运行中 - success:成功 - failed:失败

	Status string `json:"status"`
	// 任务类型。

	JobType string `json:"jobType"`
	// 任务名称。

	JobName string `json:"jobName"`
	// 关联图名称。

	RelatedGraph string `json:"relatedGraph"`
	// 任务开始时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"

	BeginTime *string `json:"beginTime,omitempty"`
	// 任务结束时间，格式为UTC,\"yyyy-MM-dd'T'HH:mm:ss\"

	EndTime *string `json:"endTime,omitempty"`

	JobDetail *JobDetail `json:"jobDetail,omitempty"`
	// 任务失败原因

	FailReason *string `json:"failReason,omitempty"`
	// 任务执行进度，预留字段，暂未使用。

	JobProgress *float64 `json:"jobProgress,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
