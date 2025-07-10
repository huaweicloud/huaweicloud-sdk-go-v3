package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTicketOperationHistoriesRequest Request Object
type ListTicketOperationHistoriesRequest struct {

	// 工单类型
	TicketType string `json:"ticket_type"`

	Body *ListTicketParams `json:"body,omitempty"`
}

func (o ListTicketOperationHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTicketOperationHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListTicketOperationHistoriesRequest", string(data)}, " ")
}
