package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSmartChatJobResponse Response Object
type StopSmartChatJobResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopSmartChatJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSmartChatJobResponse struct{}"
	}

	return strings.Join([]string{"StopSmartChatJobResponse", string(data)}, " ")
}
