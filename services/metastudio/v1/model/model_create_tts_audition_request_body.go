package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtsAuditionRequestBody 文本转语音试听请求。
type CreateTtsAuditionRequestBody struct {

	// 待合成文本。
	Text string `json:"text"`

	// 发送给tts的待合成文本。
	TtsText *string `json:"tts_text,omitempty"`

	// 音色ID，获取方式详见[获取音色ID](metastudio_02_0054.xml)。
	Emotion string `json:"emotion"`

	// 语速。 * 当取值为“100”时，表示一个成年人正常的语速，约为250字/分钟。 * 50表示0.5倍语速，100表示正常语速，200表示2倍语速。
	Speed *int32 `json:"speed,omitempty"`

	// 音高。
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。
	Volume *int32 `json:"volume,omitempty"`

	// 业务场景，多个入口调用试听接口时的业务场景
	BusinessType *string `json:"business_type,omitempty"`

	// 风格情感
	Style *string `json:"style,omitempty"`
}

func (o CreateTtsAuditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsAuditionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTtsAuditionRequestBody", string(data)}, " ")
}
