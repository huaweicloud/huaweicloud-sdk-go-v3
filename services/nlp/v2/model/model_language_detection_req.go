package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type LanguageDetectionReq struct {
	// 需要识别语种的文本。仅支持utf-8编码，长度不超过1000字符。

	Text string `json:"text"`
}

func (o LanguageDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LanguageDetectionReq struct{}"
	}

	return strings.Join([]string{"LanguageDetectionReq", string(data)}, " ")
}
