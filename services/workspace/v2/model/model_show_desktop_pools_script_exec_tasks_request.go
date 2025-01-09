package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopPoolsScriptExecTasksRequest Request Object
type ShowDesktopPoolsScriptExecTasksRequest struct {

	// 桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 脚本id。
	ScriptId *string `json:"script_id,omitempty"`

	// 脚本名称。
	ScriptName *string `json:"script_name,omitempty"`

	// 执行情况。SUCCESS：成功，FAILED：失败，RUNNING：执行中，WAITING：等待。
	Status *string `json:"status,omitempty"`

	// 查询的任务类型。支持SCRIPT、COMMAND。
	TaskType *string `json:"task_type,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。取值范围0-100，默认值是10。
	Limit *int32 `json:"limit,omitempty"`

	// 按执行时间查询的起始时间。指定该参数后，返回的结果为此时间之后的所有任务记录。时间格式如：“2021-10-01T12:00:00Z”。
	ExecuteTimeStart *string `json:"execute_time_start,omitempty"`

	// 按执行时间查询的终止时间。指定该参数后，返回的结果为此时间之前的所有任务记录。时间格式如：“2021-10-01T12:00:00Z”。
	ExecuteTimeEnd *string `json:"execute_time_end,omitempty"`

	// 任务id。
	TaskId *[]string `json:"task_id,omitempty"`
}

func (o ShowDesktopPoolsScriptExecTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopPoolsScriptExecTasksRequest struct{}"
	}

	return strings.Join([]string{"ShowDesktopPoolsScriptExecTasksRequest", string(data)}, " ")
}
