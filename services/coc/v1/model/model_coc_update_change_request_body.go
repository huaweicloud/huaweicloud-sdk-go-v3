package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocUpdateChangeRequestBody 更新变更单工单请求body体。
type CocUpdateChangeRequestBody struct {
	TicketInfo *CocUpdateChangeRequestBodyTicketInfo `json:"ticket_info,omitempty"`

	// 变更子单信息。
	SubTickets *[]CocUpdateChangeRequestBodySubTickets `json:"sub_tickets,omitempty"`

	HistoryInfo *CocUpdateChangeRequestBodyHistoryInfo `json:"history_info,omitempty"`
}

func (o CocUpdateChangeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBody struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBody", string(data)}, " ")
}
