package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocBotAnswers
type DocBotAnswers struct {

	// 答案。
	Answer string `json:"answer"`

	// 置信度。
	Score float64 `json:"score"`

	// 问题。
	Question string `json:"question"`

	AnswerDetail *DocQueryAnswerDetail `json:"answer_detail,omitempty"`

	// 候选答案列表
	Details *[]DocQueryAnswerDetail `json:"details,omitempty"`
}

func (o DocBotAnswers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocBotAnswers struct{}"
	}

	return strings.Join([]string{"DocBotAnswers", string(data)}, " ")
}
