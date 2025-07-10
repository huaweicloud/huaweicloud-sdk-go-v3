package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptRecordSimpleInfo 脚本执行记录信息。
type ScriptRecordSimpleInfo struct {

	// 脚本执行记录ID。
	Id *string `json:"id,omitempty"`

	// 脚本ID。
	ScriptId *string `json:"script_id,omitempty"`

	// 脚本名称。
	ScriptName *string `json:"script_name,omitempty"`

	// 脚本执行的任务ID。
	ScriptTaskId *string `json:"script_task_id,omitempty"`

	// 执行脚本的资源ID，如桌面ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 执行脚本的资源名称，如桌面名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型，如桌面(DESKTOP)。
	ResourceType *string `json:"resource_type,omitempty"`

	// 脚本执行开始时间。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 脚本执行结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 脚本执行状态。
	Status *string `json:"status,omitempty"`

	// 执行批次（默认：0，灰度：1，非灰度：2）。
	ExecuteOrder *int32 `json:"execute_order,omitempty"`

	// 错误码。
	ResultCode *string `json:"result_code,omitempty"`

	// 原因。
	Reason *string `json:"reason,omitempty"`
}

func (o ScriptRecordSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptRecordSimpleInfo struct{}"
	}

	return strings.Join([]string{"ScriptRecordSimpleInfo", string(data)}, " ")
}
