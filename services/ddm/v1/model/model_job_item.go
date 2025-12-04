package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobItem struct {

	// 任务id。
	Id *string `json:"id,omitempty"`

	// 任务名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 开始时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 过程。
	Process *string `json:"process,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 操作。
	Jobs *[]string `json:"jobs,omitempty"`

	// 逻辑库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o JobItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobItem struct{}"
	}

	return strings.Join([]string{"JobItem", string(data)}, " ")
}
