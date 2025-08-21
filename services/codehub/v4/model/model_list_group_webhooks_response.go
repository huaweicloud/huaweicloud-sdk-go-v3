package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupWebhooksResponse Response Object
type ListGroupWebhooksResponse struct {
	Body *[]WebhookDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGroupWebhooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupWebhooksResponse struct{}"
	}

	return strings.Join([]string{"ListGroupWebhooksResponse", string(data)}, " ")
}
