package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeedbackOptionResponse Response Object
type ListFeedbackOptionResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 选项列表
	FeedbackOptions *[]FeedbackOption `json:"feedback_options,omitempty"`
	HttpStatusCode  int               `json:"-"`
}

func (o ListFeedbackOptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeedbackOptionResponse struct{}"
	}

	return strings.Join([]string{"ListFeedbackOptionResponse", string(data)}, " ")
}
