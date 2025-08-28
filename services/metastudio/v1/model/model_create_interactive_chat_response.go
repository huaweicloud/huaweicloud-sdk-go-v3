package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInteractiveChatResponse Response Object
type CreateInteractiveChatResponse struct {

	// 答复
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInteractiveChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInteractiveChatResponse struct{}"
	}

	return strings.Join([]string{"CreateInteractiveChatResponse", string(data)}, " ")
}
