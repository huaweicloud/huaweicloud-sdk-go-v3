package model

import (
	"encoding/json"

	"strings"
)

//
type KbqaAnswers struct {
	// 答案。

	Answer string `json:"answer"`
	// 答案评分。

	Score float64 `json:"score"`
}

func (o KbqaAnswers) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KbqaAnswers struct{}"
	}

	return strings.Join([]string{"KbqaAnswers", string(data)}, " ")
}
