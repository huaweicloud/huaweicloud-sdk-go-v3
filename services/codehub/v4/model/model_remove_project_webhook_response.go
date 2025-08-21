package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveProjectWebhookResponse Response Object
type RemoveProjectWebhookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveProjectWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveProjectWebhookResponse struct{}"
	}

	return strings.Join([]string{"RemoveProjectWebhookResponse", string(data)}, " ")
}
