package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTicketInfoRequest Request Object
type DeleteTicketInfoRequest struct {

	// 需要操作的工单类型，需要传值change表示需要删除的工单为变更单。
	TicketType string `json:"ticket_type"`

	// 需要删除的工单单号。
	TicketId string `json:"ticket_id"`
}

func (o DeleteTicketInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTicketInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteTicketInfoRequest", string(data)}, " ")
}
