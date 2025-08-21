package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveRepositoryWebhookResponse Response Object
type RemoveRepositoryWebhookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveRepositoryWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveRepositoryWebhookResponse struct{}"
	}

	return strings.Join([]string{"RemoveRepositoryWebhookResponse", string(data)}, " ")
}
