package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptTasksRequest Request Object
type ListScriptTasksRequest struct {

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 执行脚本的资源组ID。
	ResourceGroupId *[]string `json:"resource_group_id,omitempty"`

	// 脚本ID。
	ScriptId *string `json:"script_id,omitempty"`

	// 脚本名。
	ScriptName *string `json:"script_name,omitempty"`

	// 执行情况。
	Status *string `json:"status,omitempty"`

	// 资源组类型。
	ResourceGroupType *string `json:"resource_group_type,omitempty"`

	// 执行脚本的任务ID。
	TaskId *[]string `json:"task_id,omitempty"`

	// 任务类型(SCRIPT/COMMAND)。
	TaskType *string `json:"task_type,omitempty"`

	// 按执行时间查询的起始时间。指定该参数后，返回的结果为此时间之后的所有任务记录。时间格式如：“2021-10-01T12:00:00Z”。
	ExecuteTimeStart *string `json:"execute_time_start,omitempty"`

	// 按执行时间查询的终止时间。指定该参数后，返回的结果为此时间之前的所有任务记录。时间格式如：“2021-10-01T12:00:00Z”。
	ExecuteTimeEnd *string `json:"execute_time_end,omitempty"`
}

func (o ListScriptTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScriptTasksRequest", string(data)}, " ")
}
