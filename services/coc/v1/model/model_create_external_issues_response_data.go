package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalIssuesResponseData CreateExternalIssuesResponseData
type CreateExternalIssuesResponseData struct {

	// 问题单号
	TicketId string `json:"ticket_id"`
}

func (o CreateExternalIssuesResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalIssuesResponseData struct{}"
	}

	return strings.Join([]string{"CreateExternalIssuesResponseData", string(data)}, " ")
}
