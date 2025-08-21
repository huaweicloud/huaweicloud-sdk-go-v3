package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectWebhooksResponse Response Object
type ListProjectWebhooksResponse struct {
	Body *[]WebhookDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectWebhooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectWebhooksResponse struct{}"
	}

	return strings.Join([]string{"ListProjectWebhooksResponse", string(data)}, " ")
}
