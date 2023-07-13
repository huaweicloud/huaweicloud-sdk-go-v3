package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VmsSendTask 智能信息基础版任务详情。
type VmsSendTask struct {

	// 智能信息基础版任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息基础版任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 智能信息基础版任务状态。
	TaskState *string `json:"task_state,omitempty"`

	// 智能信息基础版模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息基础版模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 操作员。
	Operator *string `json:"operator,omitempty"`

	// 智能信息基础版任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 智能信息基础版任务发送时间。
	SendTime *string `json:"send_time,omitempty"`

	// 需要发送的手机号码总数（有效号码总数）。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 已经发送成功的手机号码总数。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 已经发送失败的手机号码总数。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 已经发送超时失败的手机号码总数（72小时没接收到状态认定为超时失败）。
	TimeoutCount *int32 `json:"timeout_count,omitempty"`

	// 扩展字段。  > 预留字段。
	Exdata *string `json:"exdata,omitempty"`
}

func (o VmsSendTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VmsSendTask struct{}"
	}

	return strings.Join([]string{"VmsSendTask", string(data)}, " ")
}
