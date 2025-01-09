package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptRecordsRequest Request Object
type ListScriptRecordsRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 执行脚本的资源ID列表。
	ResourceId *[]string `json:"resource_id,omitempty"`

	// 执行脚本的资源组ID。
	ResourceGroupId *[]string `json:"resource_group_id,omitempty"`

	// 执行的脚本ID。
	ScriptId *[]string `json:"script_id,omitempty"`

	// 执行的脚本名称。
	ScriptName *[]string `json:"script_name,omitempty"`

	// 执行脚本的执行情况。
	Status *string `json:"status,omitempty"`

	// 是否首批执行。
	IsFirstOrder *bool `json:"is_first_order,omitempty"`

	// 执行脚本的任务ID。
	ScriptTaskId *string `json:"script_task_id,omitempty"`

	// 执行记录的任务类型(SCRIPT/COMMAND)。
	TaskType *string `json:"task_type,omitempty"`

	// 是否查询历史记录，默认为false，为true时需要同时传入resource_id与script_id。
	ShowHistory *bool `json:"show_history,omitempty"`

	// 按执行时间查询的起始时间。指定该参数后，返回的结果为此时间之后的所有执行记录。时间格式如：“2021-10-01T12:00:00Z”。
	ExecuteTimeStart *string `json:"execute_time_start,omitempty"`

	// 按执行时间查询的终止时间。指定该参数后，返回的结果为此时间之前的所有执行记录。时间格式如：“2021-10-01T12:00:00Z”。
	ExecuteTimeEnd *string `json:"execute_time_end,omitempty"`
}

func (o ListScriptRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptRecordsRequest", string(data)}, " ")
}
