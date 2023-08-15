package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunModerationAudioResponseBodyResultDetailAudio 返回的音频审核结果列表，porn为涉黄场景审核结果； 当前仅支持porn场景。
type RunModerationAudioResponseBodyResultDetailAudio struct {

	// 涉黄场景审核结果
	Porn *[]PornModerationResultDetail `json:"porn,omitempty"`
}

func (o RunModerationAudioResponseBodyResultDetailAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioResponseBodyResultDetailAudio struct{}"
	}

	return strings.Join([]string{"RunModerationAudioResponseBodyResultDetailAudio", string(data)}, " ")
}
