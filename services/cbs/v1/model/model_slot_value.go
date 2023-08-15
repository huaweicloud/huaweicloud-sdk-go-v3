package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlotValue
type SlotValue struct {

	// 词。
	Word string `json:"word"`

	// 归一化后的标准词。
	NormWord string `json:"norm_word"`

	// 词的起始位置。
	BeginPosition int32 `json:"begin_position"`

	// 词的结束位置。
	EndPosition int32 `json:"end_position"`
}

func (o SlotValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlotValue struct{}"
	}

	return strings.Join([]string{"SlotValue", string(data)}, " ")
}
