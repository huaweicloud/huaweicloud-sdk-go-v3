package model

import (
	"encoding/json"

	"strings"
)

//
type HistorySlotWord struct {
	// 词

	Word string `json:"word"`
	// 归一化后的词

	NormWord string `json:"norm_word"`
}

func (o HistorySlotWord) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HistorySlotWord struct{}"
	}

	return strings.Join([]string{"HistorySlotWord", string(data)}, " ")
}
