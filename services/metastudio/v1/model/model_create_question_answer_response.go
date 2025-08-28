package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQuestionAnswerResponse Response Object
type CreateQuestionAnswerResponse struct {

	// 问答对ID。
	QuestionAnswerId *string `json:"question_answer_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQuestionAnswerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQuestionAnswerResponse struct{}"
	}

	return strings.Join([]string{"CreateQuestionAnswerResponse", string(data)}, " ")
}
