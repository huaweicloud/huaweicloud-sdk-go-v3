package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTicketsRequest Request Object
type ListTicketsRequest struct {

	// 工单类型
	TicketType string `json:"ticket_type"`

	Body *ListTicketParamsWithPage `json:"body,omitempty"`
}

func (o ListTicketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTicketsRequest struct{}"
	}

	return strings.Join([]string{"ListTicketsRequest", string(data)}, " ")
}
