package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceWebhookResponse Response Object
type CreateInstanceWebhookResponse struct {

	// 触发器ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInstanceWebhookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceWebhookResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceWebhookResponse", string(data)}, " ")
}
