package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryWebhooksResponse Response Object
type ListRepositoryWebhooksResponse struct {
	Body *[]WebhookDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryWebhooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryWebhooksResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryWebhooksResponse", string(data)}, " ")
}
