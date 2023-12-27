package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobDetail 任务信息。
type JobDetail struct {

	// 任务ID
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 任务执行状态。 取值： Running：表示任务正在执行。 Completed：表示任务执行成功。 Failed：表示任务执行失败。  枚举值： Running Completed Failed
	Status string `json:"status"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量[，例如北京时间偏移显示为+0800。](tag:hc)
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量[，例如北京时间偏移显示为+0800。](tag:hc)
	EndTime string `json:"end_time"`

	// 任务执行进度。 > 执行中状态才返回执行进度，例如“60%”，表示任务执行进度为60%，否则返回“”。
	Progress string `json:"progress"`

	Instance *JobInstanceInfo `json:"instance"`

	// 任务执行失败时的错误信息。
	FailReason string `json:"fail_reason"`
}

func (o JobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}
