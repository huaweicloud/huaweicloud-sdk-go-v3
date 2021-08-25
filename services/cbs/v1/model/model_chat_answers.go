package model

import (
	"encoding/json"

	"strings"
)

//
type ChatAnswers struct {
	// 答案。

	Answer string `json:"answer"`
	// 闲聊的置信度，范围:0.0~1.0  0.0表示兜底回复

	Score *float32 `json:"score,omitempty"`
}

func (o ChatAnswers) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChatAnswers struct{}"
	}

	return strings.Join([]string{"ChatAnswers", string(data)}, " ")
}
