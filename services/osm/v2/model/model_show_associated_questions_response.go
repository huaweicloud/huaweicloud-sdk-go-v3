package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAssociatedQuestionsResponse struct {

	// 问题列表
	Questions *[]string `json:"questions,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssociatedQuestionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssociatedQuestionsResponse struct{}"
	}

	return strings.Join([]string{"ShowAssociatedQuestionsResponse", string(data)}, " ")
}
