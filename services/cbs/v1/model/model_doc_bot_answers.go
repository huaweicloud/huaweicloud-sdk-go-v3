package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DocBotAnswers struct {

	// 答案。
	Answer string `json:"answer" xml:"answer"`

	// 置信度。
	Score float64 `json:"score" xml:"score"`

	// 问题。
	Question string `json:"question" xml:"question"`

	AnswerDetail *DocQueryAnswerDetail `json:"answer_detail,omitempty" xml:"answer_detail"`

	// 候选答案列表
	Details *[]DocQueryAnswerDetail `json:"details,omitempty" xml:"details"`
}

func (o DocBotAnswers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocBotAnswers struct{}"
	}

	return strings.Join([]string{"DocBotAnswers", string(data)}, " ")
}
