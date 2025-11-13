package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMessageFeedbackResponse Response Object
type CreateMessageFeedbackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMessageFeedbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageFeedbackResponse struct{}"
	}

	return strings.Join([]string{"CreateMessageFeedbackResponse", string(data)}, " ")
}
