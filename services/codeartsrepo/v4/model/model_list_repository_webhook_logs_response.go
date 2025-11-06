package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryWebhookLogsResponse Response Object
type ListRepositoryWebhookLogsResponse struct {
	Body *[]WebhookLogDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryWebhookLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryWebhookLogsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryWebhookLogsResponse", string(data)}, " ")
}
