package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListScheduleJobsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0，必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数。默认为10
	Limit *string `json:"limit,omitempty"`

	// 任务执行状态。 取值： - 值为“Running”，表示任务正在执行。 - 值为“Completed”，表示任务执行成功。 - 值为“Failed”，表示任务执行失败。 - 值为“Pending”，表示任务未执行。
	Status *string `json:"status,omitempty"`

	// 起始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。 说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。 说明：创建时返回值为空，数据库实例创建成功后该值不为空。
	EndTime *string `json:"end_time,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务调度类型。
	JobName *string `json:"job_name,omitempty"`
}

func (o ListScheduleJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleJobsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduleJobsRequest", string(data)}, " ")
}
