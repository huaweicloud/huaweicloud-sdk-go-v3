package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulHandleTaskResponseInfo 漏洞处置任务信息
type VulHandleTaskResponseInfo struct {

	// **参数解释**: 任务ID **取值范围**: 字符长度1-64位
	Id *string `json:"id,omitempty"`

	// **参数解释**: 处置任务类型 **取值范围**: - immediate_repair：立即修复。 - verify：验证。
	HandleType *string `json:"handle_type,omitempty"`

	// **参数解释**: 任务状态 **取值范围**: - running：执行中。 - finished：已完成。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 任务开始时间 **取值范围**: 最小值0，最大值9223372036854775807
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**: 任务结束时间 **取值范围**: 最小值0，最大值9223372036854775807
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**: 任务涉及的漏洞类型列表 **取值范围**: 最小值0，最大值100
	VulTypes *[]string `json:"vul_types,omitempty"`

	// **参数解释**: 等待执行的数量 **取值范围**: 最小值0，最大值2147483647
	ReadyNum *int32 `json:"ready_num,omitempty"`

	// **参数解释**: 执行中的数量 **取值范围**: 最小值0，最大值2147483647
	RunningNum *int32 `json:"running_num,omitempty"`

	// **参数解释**: 执行成功的数量 **取值范围**: 最小值0，最大值2147483647
	SuccessNum *int32 `json:"success_num,omitempty"`

	// **参数解释**: 执行失败的数量 **取值范围**: 最小值0，最大值2147483647
	FailedNum *int32 `json:"failed_num,omitempty"`

	// **参数解释**: 预计还需要多长时间 **取值范围**: 最小值0，最大值604800
	EstimatedTime *float32 `json:"estimated_time,omitempty"`
}

func (o VulHandleTaskResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHandleTaskResponseInfo struct{}"
	}

	return strings.Join([]string{"VulHandleTaskResponseInfo", string(data)}, " ")
}
