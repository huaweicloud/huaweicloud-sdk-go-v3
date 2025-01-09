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

	// 任务执行状态。  取值： 值为“Running”，表示任务正在执行。 值为“Completed”，表示任务执行成功。 值为“Failed”，表示任务执行失败。
	Status string `json:"status"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedTime string `json:"created_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime string `json:"end_time"`

	// 任务执行进度。  执行中状态才返回执行进度，例如“60%”，表示任务执行进度为60%，否则返回“”。
	Process *string `json:"process,omitempty"`

	Instance *Instance `json:"instance"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}
