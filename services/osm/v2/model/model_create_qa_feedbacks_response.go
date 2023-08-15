package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQaFeedbacksResponse Response Object
type CreateQaFeedbacksResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 反馈记录id
	FeedbackId     *string `json:"feedback_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQaFeedbacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQaFeedbacksResponse struct{}"
	}

	return strings.Join([]string{"CreateQaFeedbacksResponse", string(data)}, " ")
}
