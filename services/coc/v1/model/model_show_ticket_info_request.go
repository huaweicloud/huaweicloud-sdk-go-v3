package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTicketInfoRequest Request Object
type ShowTicketInfoRequest struct {

	// 工单类型
	TicketType string `json:"ticket_type"`

	// ID of Ticket
	TicketId string `json:"ticket_id"`
}

func (o ShowTicketInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTicketInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowTicketInfoRequest", string(data)}, " ")
}
