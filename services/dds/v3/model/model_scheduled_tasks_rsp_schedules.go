package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduledTasksRspSchedules struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务名称。对应取值如下：   \"RESIZE_FLAVOR\"：变更实例的CPU和内存规格
	JobName *string `json:"job_name,omitempty"`

	// 任务执行状态。 取值：  值为\"Running\"，表示任务正在执行。  值为\"Completed\"，表示任务执行成功。  值为\"Failed\"，表示任务执行失败。  值为\"Pending\"，表示任务未执行。
	JobStatus *string `json:"job_status,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态。 取值：  值为“createfail”，表示实例创建失败。  值为“creating”，表示实例创建中。  值为“normal”，表示实例正常。  值为“abnormal”，表示实例异常。  值为“deleted”，表示实例已删除。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 任务创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务开始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ScheduledTasksRspSchedules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTasksRspSchedules struct{}"
	}

	return strings.Join([]string{"ScheduledTasksRspSchedules", string(data)}, " ")
}
