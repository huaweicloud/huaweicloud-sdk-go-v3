package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobInfo struct {

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 任务状态。取值为“Running”为执行中； 取值为“Completed”为完成； 取值为“Failed” 为失败。
	Status string `json:"status"`

	// 任务执行进度。 说明 执行中状态才返回执行进度，例如“60%”，表示任务执行进度为60%，否则返回“”。 任务执行进度。
	Progress string `json:"progress"`

	// 任务执行失败时的错误信息。
	FailReason string `json:"fail_reason"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt string `json:"created_at"`

	// 结束时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndedAt string `json:"ended_at"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}
