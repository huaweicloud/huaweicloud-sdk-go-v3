package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobDetail struct {

	// 任务ID
	Id string `json:"id" xml:"id"`

	// 任务名称。
	Name string `json:"name" xml:"name"`

	// 任务执行状态。
	Status string `json:"status" xml:"status"`

	// 任务创建时间，格式为yyyy-mm-ddThh:mm:ssZ。
	Created string `json:"created" xml:"created"`

	// 任务结束时间，格式为yyyy-mm-ddThh:mm:ssZ。
	Ended string `json:"ended" xml:"ended"`

	// 任务执行进度。
	Progress string `json:"progress" xml:"progress"`

	Instance *JobInstanceInfo `json:"instance" xml:"instance"`

	// 任务执行失败时的错误信息。
	FailReason string `json:"fail_reason" xml:"fail_reason"`
}

func (o JobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}
