package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocUpdateChangeRequestBodyV2 UpdateChangeRequestBodyV2 更新变更单
type CocUpdateChangeRequestBodyV2 struct {
	TicketInfo *CocUpdateChangeRequestBodyV2TicketInfo `json:"ticket_info,omitempty"`

	// 子单信息
	SubTickets *[]CocUpdateChangeRequestBodyV2SubTickets `json:"sub_tickets,omitempty"`

	HistoryInfo *CocUpdateChangeRequestBodyV2HistoryInfo `json:"history_info,omitempty"`
}

func (o CocUpdateChangeRequestBodyV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBodyV2 struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBodyV2", string(data)}, " ")
}
