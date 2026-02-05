package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenerateSpeechRequestBody struct {

	// 待合成的文本，文本长度不大于300字符。
	Text string `json:"text"`

	Config *GenerateSpeechRequestBodyConfig `json:"config"`
}

func (o GenerateSpeechRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSpeechRequestBody struct{}"
	}

	return strings.Join([]string{"GenerateSpeechRequestBody", string(data)}, " ")
}
