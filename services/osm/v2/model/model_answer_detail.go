package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnswerDetail struct {

	// 答案列表
	QabotAnswers *[]QabotAnswer `json:"qabot_answers,omitempty"`

	QaFlowAnswers *QaFlowHitResult `json:"qa_flow_answers,omitempty"`

	// 问题
	ChatAnswer *string `json:"chat_answer,omitempty"`

	GraphAnswer *QaGraphAnswer `json:"graph_answer,omitempty"`
}

func (o AnswerDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnswerDetail struct{}"
	}

	return strings.Join([]string{"AnswerDetail", string(data)}, " ")
}
