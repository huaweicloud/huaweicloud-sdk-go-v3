package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveGroupWebhookResponse Response Object
type RemoveGroupWebhookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveGroupWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveGroupWebhookResponse struct{}"
	}

	return strings.Join([]string{"RemoveGroupWebhookResponse", string(data)}, " ")
}
