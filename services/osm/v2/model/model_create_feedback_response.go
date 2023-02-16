package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFeedbackResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFeedbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeedbackResponse struct{}"
	}

	return strings.Join([]string{"CreateFeedbackResponse", string(data)}, " ")
}
