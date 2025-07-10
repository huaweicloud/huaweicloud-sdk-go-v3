package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptTaskInfo 脚本任务信息。
type ScriptTaskInfo struct {

	// 脚本任务ID。
	Id *string `json:"id,omitempty"`

	// 脚本列表。
	TaskScripts *[]ScriptTaskInfoTaskScripts `json:"task_scripts,omitempty"`

	// 资源类型，如DESKTOP。
	ResourceType *string `json:"resource_type,omitempty"`

	// 执行脚本的资源ID列表。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 灰度批次执行资源ID列表。
	GrayResourceIds *[]string `json:"gray_resource_ids,omitempty"`

	// task中成功的执行记录数量。
	SuccessNum *int32 `json:"success_num,omitempty"`

	// task中失败的执行记录数量。
	FailedNum *int32 `json:"failed_num,omitempty"`

	// task中跳过的执行记录数量。
	SkipNum *int32 `json:"skip_num,omitempty"`

	// 脚本执行开始时间。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 脚本执行结束时间。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 任务结果。
	Status *string `json:"status,omitempty"`
}

func (o ScriptTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptTaskInfo struct{}"
	}

	return strings.Join([]string{"ScriptTaskInfo", string(data)}, " ")
}
