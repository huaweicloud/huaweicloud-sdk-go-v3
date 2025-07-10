package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTicketActionRequest Request Object
type ExecuteTicketActionRequest struct {

	// 工单类型，变更单传值change，问题单传值issues_mgmt。
	TicketType string `json:"ticket_type"`

	Body *ExecuteActionParams `json:"body,omitempty"`
}

func (o ExecuteTicketActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTicketActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTicketActionRequest", string(data)}, " ")
}
