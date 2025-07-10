package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteActionParams struct {

	// 执行当前操作的工单单号。
	TicketId *string `json:"ticket_id,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 动作名称
	Action *string `json:"action,omitempty"`

	// 当前执行的动作信息
	Params *interface{} `json:"params,omitempty"`
}

func (o ExecuteActionParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteActionParams struct{}"
	}

	return strings.Join([]string{"ExecuteActionParams", string(data)}, " ")
}
