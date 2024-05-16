package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtsAuditionRequestBody 文本转语音试听请求。
type CreateTtsAuditionRequestBody struct {

	// 待合成文本。
	Text string `json:"text"`

	// 音色ID。
	Emotion string `json:"emotion"`

	// 语速。 默认值100，最小值50，最大值200。 > * 当取值为“100”时，表示一个成年人正常的语速，约为250字/分钟。 > * 50表示0.5倍语速，100表示正常语速，200表示2倍语速。
	Speed *int32 `json:"speed,omitempty"`

	// 音高。 默认值100，最小值50，最大值200。
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。 默认值140，最小值90，最大值240。
	Volume *int32 `json:"volume,omitempty"`
}

func (o CreateTtsAuditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsAuditionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTtsAuditionRequestBody", string(data)}, " ")
}
