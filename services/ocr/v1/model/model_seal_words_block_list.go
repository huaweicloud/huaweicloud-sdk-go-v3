package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SealWordsBlockList struct {

	// 印章文本块。
	Words *string `json:"words,omitempty"`

	// 印章文本块的置信度。
	WordsConfidence *float32 `json:"words_confidence,omitempty"`
}

func (o SealWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SealWordsBlockList struct{}"
	}

	return strings.Join([]string{"SealWordsBlockList", string(data)}, " ")
}
