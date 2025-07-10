package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTicketRequest Request Object
type UpdateTicketRequest struct {

	// 需要修改的工单类型，此处需传递固定值change，表示更新变更单状态。
	TicketType string `json:"ticket_type"`

	// 变更单工单单号。
	TicketId string `json:"ticket_id"`

	Body *CocUpdateChangeRequestBody `json:"body,omitempty"`
}

func (o UpdateTicketRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTicketRequest struct{}"
	}

	return strings.Join([]string{"UpdateTicketRequest", string(data)}, " ")
}
