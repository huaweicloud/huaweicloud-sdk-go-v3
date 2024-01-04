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

	// 语速。
	Speed *int32 `json:"speed,omitempty"`

	// 音高。
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。
	Volume *int32 `json:"volume,omitempty"`
}

func (o CreateTtsAuditionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsAuditionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTtsAuditionRequestBody", string(data)}, " ")
}
