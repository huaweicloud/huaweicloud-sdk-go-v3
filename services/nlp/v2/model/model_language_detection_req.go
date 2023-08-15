package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LanguageDetectionReq This is a auto create Body Object
type LanguageDetectionReq struct {

	// 需要识别语种的文本。仅支持utf-8编码，长度不超过2000字符。
	Text string `json:"text"`
}

func (o LanguageDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LanguageDetectionReq struct{}"
	}

	return strings.Join([]string{"LanguageDetectionReq", string(data)}, " ")
}
