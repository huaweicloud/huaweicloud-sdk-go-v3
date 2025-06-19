package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWebhookUrlResponse Response Object
type CheckWebhookUrlResponse struct {

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckWebhookUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWebhookUrlResponse struct{}"
	}

	return strings.Join([]string{"CheckWebhookUrlResponse", string(data)}, " ")
}
