package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuestionAnswerResponse Response Object
type ShowQuestionAnswerResponse struct {

	// 问答对ID。
	QuestionAnswerId *string `json:"question_answer_id,omitempty"`

	// 标准问题。
	Question *string `json:"question,omitempty"`

	// 问题答案。
	Answer *string `json:"answer,omitempty"`

	// 所有相似问题，多个相似问题间用换行符\\n分隔。
	SimilarQuestions *string `json:"similar_questions,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQuestionAnswerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuestionAnswerResponse struct{}"
	}

	return strings.Join([]string{"ShowQuestionAnswerResponse", string(data)}, " ")
}
