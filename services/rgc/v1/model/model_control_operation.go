package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ControlOperation 控制策略实施的情况。
type ControlOperation struct {

	// 本次操作控制策略的ID。
	OperationControlStatusId *string `json:"operation_control_status_id,omitempty"`

	// 操作类型，启用控制策略或禁用控制策略。
	OperationType *string `json:"operation_type,omitempty"`

	// 控制策略实施的状态 SUCCEEDED | FAILED | IN_PROGRESS。
	Status *string `json:"status,omitempty"`

	// 控制策略实施失败的错误信息。
	Message *string `json:"message,omitempty"`

	// 操作开始的时间。
	StartTime *string `json:"start_time,omitempty"`

	// 操作结束的时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ControlOperation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlOperation struct{}"
	}

	return strings.Join([]string{"ControlOperation", string(data)}, " ")
}
