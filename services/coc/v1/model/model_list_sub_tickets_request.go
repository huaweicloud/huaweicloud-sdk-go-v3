package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubTicketsRequest Request Object
type ListSubTicketsRequest struct {

	// 工单类型，此处传固定值change。
	TicketType string `json:"ticket_type"`

	// 变更单工单id
	TicketId string `json:"ticket_id"`

	// 资源类型
	Type string `json:"type"`

	// 每页显示的条数
	Limit *int64 `json:"limit,omitempty"`

	// 上一页数据的最后一条记录的id
	Marker *string `json:"marker,omitempty"`
}

func (o ListSubTicketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubTicketsRequest struct{}"
	}

	return strings.Join([]string{"ListSubTicketsRequest", string(data)}, " ")
}
