package model

import (
	"encoding/json"

	"strings"
)

//
type MathInfo struct {
	// 数学试卷返回结果，表示题号。

	QuestionNumber *string `json:"question_number,omitempty"`
	// 数学试卷答案的文字块数目。

	AnswerBlockCount *int32 `json:"answer_block_count,omitempty"`
	// 数学试卷答案识别文字块列表，输出顺序从左到右，从上到下。

	AnswerBlockList *[]AnswerBlockList `json:"answer_block_list,omitempty"`
}

func (o MathInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MathInfo struct{}"
	}

	return strings.Join([]string{"MathInfo", string(data)}, " ")
}
