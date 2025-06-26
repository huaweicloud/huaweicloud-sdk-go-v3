package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceWebhookResponse Response Object
type DeleteInstanceWebhookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceWebhookResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceWebhookResponse", string(data)}, " ")
}
