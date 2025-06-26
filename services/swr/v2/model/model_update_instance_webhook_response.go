package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceWebhookResponse Response Object
type UpdateInstanceWebhookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceWebhookResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceWebhookResponse", string(data)}, " ")
}
