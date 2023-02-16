package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateQuestionInSessionResponse struct {

	// 答案列表
	Answers *[]AnswerQaPair `json:"answers,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 答案类型：0 知识库回复,1 技能回复,2 闲聊回复,3 华为云答案。
	AnswerType *int32 `json:"answer_type,omitempty"`

	// 请求Id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQuestionInSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQuestionInSessionResponse struct{}"
	}

	return strings.Join([]string{"CreateQuestionInSessionResponse", string(data)}, " ")
}
