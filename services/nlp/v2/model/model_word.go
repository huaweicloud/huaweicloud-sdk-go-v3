package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Word 词汇类
type Word struct {

	// 词汇文本。
	Content string `json:"content"`

	// 词汇对应的词性。
	Pos string `json:"pos"`
}

func (o Word) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Word struct{}"
	}

	return strings.Join([]string{"Word", string(data)}, " ")
}
