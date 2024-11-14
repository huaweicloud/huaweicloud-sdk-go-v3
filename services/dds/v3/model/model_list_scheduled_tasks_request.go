package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRequest Request Object
type ListScheduledTasksRequest struct {

	// 任务名称，默认为空。对应取值如下：  \"RESIZE_FLAVOR\"：变更实例的CPU和内存规格
	JobName *string `json:"job_name,omitempty"`

	// 任务执行状态，默认为空。 取值：  值为\"Pending\"，表示任务未执行。  值为\"Running\"，表示任务正在执行。  值为\"Completed\"，表示任务执行成功。  值为\"Failed\"，表示任务执行失败。  值为\"Canceled\"，表示任务取消执行。
	JobStatus *string `json:"job_status,omitempty"`

	// 实例ID，不传该值默认查所有符合条件的实例。
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务创建起始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100，不传默认为当前时间前七天。
	StartTime *string `json:"start_time,omitempty"`

	// 任务创建结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100，不传默认为当前时间。
	EndTime *string `json:"end_time,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。不传该参数时，默认为10,取值范围1-100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRequest", string(data)}, " ")
}
