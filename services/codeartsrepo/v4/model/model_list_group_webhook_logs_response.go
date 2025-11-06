package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupWebhookLogsResponse Response Object
type ListGroupWebhookLogsResponse struct {
	Body *[]WebhookLogExtendDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGroupWebhookLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupWebhookLogsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupWebhookLogsResponse", string(data)}, " ")
}
