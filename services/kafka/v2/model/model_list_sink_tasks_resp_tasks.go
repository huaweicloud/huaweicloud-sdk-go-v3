package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSinkTasksRespTasks struct {

	// 任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 转储任务名称。
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 转储任务类型。
	DestinationType *string `json:"destination_type,omitempty" xml:"destination_type"`

	// 转储任务创建时间戳。
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time"`

	// 转储任务状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 返回任务转存的topics列表或者正则表达式。
	Topics *string `json:"topics,omitempty" xml:"topics"`
}

func (o ListSinkTasksRespTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSinkTasksRespTasks struct{}"
	}

	return strings.Join([]string{"ListSinkTasksRespTasks", string(data)}, " ")
}
