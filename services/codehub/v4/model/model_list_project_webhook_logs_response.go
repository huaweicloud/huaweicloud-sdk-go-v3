package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectWebhookLogsResponse Response Object
type ListProjectWebhookLogsResponse struct {
	Body *[]WebhookLogExtendDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectWebhookLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectWebhookLogsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectWebhookLogsResponse", string(data)}, " ")
}
