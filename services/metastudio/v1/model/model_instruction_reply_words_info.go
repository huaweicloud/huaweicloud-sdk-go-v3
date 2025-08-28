package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstructionReplyWordsInfo 指令回复话术信息。
type InstructionReplyWordsInfo struct {
	Language *LanguageEnum `json:"language"`

	// 回复话术，多个回复话术间用换行符\\n分隔。
	ReplyWords string `json:"reply_words"`
}

func (o InstructionReplyWordsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstructionReplyWordsInfo struct{}"
	}

	return strings.Join([]string{"InstructionReplyWordsInfo", string(data)}, " ")
}
