package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCocIssuesDetailRequest Request Object
type ShowCocIssuesDetailRequest struct {

	// 问题单号
	TicketId string `json:"ticket_id"`
}

func (o ShowCocIssuesDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCocIssuesDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCocIssuesDetailRequest", string(data)}, " ")
}
