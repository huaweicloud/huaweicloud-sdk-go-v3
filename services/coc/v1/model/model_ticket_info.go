package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TicketInfo requestBody中包含的，前端选择的四号单的信息
type TicketInfo struct {

	// 四号单id
	TicketId *string `json:"ticket_id,omitempty"`

	// 四号单类型，可选CHANGE/INCIDENT/ALARM/WARROOM
	TicketType *string `json:"ticket_type,omitempty"`

	// 四号单关联的应用id
	TargetId *string `json:"target_id,omitempty"`

	// region id
	ScopeId *string `json:"scope_id,omitempty"`
}

func (o TicketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TicketInfo struct{}"
	}

	return strings.Join([]string{"TicketInfo", string(data)}, " ")
}
