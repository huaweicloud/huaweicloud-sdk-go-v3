package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimSendTasksMode 智能信息发送任务。
type AimSendTasksMode struct {

	// 智能信息发送任务名称。
	TaskName *string `json:"task_name,omitempty"`

	SmsChannel *AimSendTaskSmsChannel `json:"sms_channel,omitempty"`

	ResolveTask *AimResolveTaskMode `json:"resolve_task,omitempty"`

	// 任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务状态。  - Success：发送成功 - Failed：发送失败  > 此状态仅代表任务提交状态，不代表智能信息发送结果。用户手机接收智能信息结果请以收到的回执结果为准，也可通过查询智能信息发送明细API获取或登录KooMessage控制台查看。
	TaskState *string `json:"task_state,omitempty"`

	// 创建时间。样例：2019-10-12T07:20:50.522Z。
	CreationTime *sdktime.SdkTime `json:"creation_time,omitempty"`

	// 提交的手机号码总数。
	SubmissionCount *int32 `json:"submission_count,omitempty"`

	// 发送数量。
	SendCount *int32 `json:"send_count,omitempty"`

	// 智能信息解析成功的手机号码总数。
	ResolveCount *int32 `json:"resolve_count,omitempty"`

	// 支持智能信息解析的手机号码总数。  >通过API发送的智能信息任务不做解析能力判断，返回-1作为标识。
	SupportResolveCount *int32 `json:"support_resolve_count,omitempty"`

	// 短链生成失败列表。  > 预留字段。
	FailedShortChains *[]CreateResolveTaskParamMode `json:"failed_short_chains,omitempty"`

	// 发送成功的智能信息数。
	SuccessNumber *int32 `json:"success_number,omitempty"`
}

func (o AimSendTasksMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimSendTasksMode struct{}"
	}

	return strings.Join([]string{"AimSendTasksMode", string(data)}, " ")
}
