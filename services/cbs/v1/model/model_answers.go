package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Answers
type Answers struct {

	// Answers.
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 标准问题。
	StQuestion *string `json:"st_question,omitempty"`

	// 相似度得分，精确到小数点后3位。
	Score *float64 `json:"score,omitempty"`

	// Domain.
	Domain *string `json:"domain,omitempty"`

	// 答案。
	Answer *string `json:"answer,omitempty"`

	// 扩展问。
	ExQuestions *[]ExQuestions `json:"ex_questions,omitempty"`
}

func (o Answers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Answers struct{}"
	}

	return strings.Join([]string{"Answers", string(data)}, " ")
}
