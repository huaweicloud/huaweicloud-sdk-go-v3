package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWebhookUrlRequest Request Object
type CheckWebhookUrlRequest struct {
	Body *CheckWebhookUrlRequestBody `json:"body,omitempty"`
}

func (o CheckWebhookUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWebhookUrlRequest struct{}"
	}

	return strings.Join([]string{"CheckWebhookUrlRequest", string(data)}, " ")
}
