package model

import (
	"encoding/json"

	"strings"
)

//
type QaBotAnswers struct {
	// 问答机器人回复。

	Answers *[]QaBotAnswer `json:"answers,omitempty"`
	// 请求ID。

	RequestId string `json:"request_id"`
}

func (o QaBotAnswers) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QaBotAnswers struct{}"
	}

	return strings.Join([]string{"QaBotAnswers", string(data)}, " ")
}
