package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQuestionAnswerReq 修改问答对请求。
type UpdateQuestionAnswerReq struct {

	// 标准问题。
	Question *string `json:"question,omitempty"`

	// 问题答案。
	Answer *string `json:"answer,omitempty"`

	// 所有相似问题，多个相似问题间用换行符\\n分隔。
	SimilarQuestions *string `json:"similar_questions,omitempty"`
}

func (o UpdateQuestionAnswerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQuestionAnswerReq struct{}"
	}

	return strings.Join([]string{"UpdateQuestionAnswerReq", string(data)}, " ")
}
