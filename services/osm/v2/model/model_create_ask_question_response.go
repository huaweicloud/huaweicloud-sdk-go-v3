package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAskQuestionResponse Response Object
type CreateAskQuestionResponse struct {

	// 问题
	Question *string `json:"question,omitempty"`

	// 答案列表
	Answers *[]AnswerQaPair `json:"answers,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 请求Id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAskQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAskQuestionResponse struct{}"
	}

	return strings.Join([]string{"CreateAskQuestionResponse", string(data)}, " ")
}
