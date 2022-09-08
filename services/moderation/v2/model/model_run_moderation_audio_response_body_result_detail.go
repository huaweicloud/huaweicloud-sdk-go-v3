package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 审核结果列表，text为语音转文本后审核结果，audio为音频审核结果。
type RunModerationAudioResponseBodyResultDetail struct {

	// 返回的语音转文本后审核结果详细信息： ● politics：涉政敏感词列表。 ● porn：涉黄敏感词列表。 ● ad：广告敏感词列表。 ● abuse：辱骂敏感词列表。 ● contraband：违禁品敏感词列表
	Text *interface{} `json:"text,omitempty"`

	Audio *RunModerationAudioResponseBodyResultDetailAudio `json:"audio,omitempty"`
}

func (o RunModerationAudioResponseBodyResultDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioResponseBodyResultDetail struct{}"
	}

	return strings.Join([]string{"RunModerationAudioResponseBodyResultDetail", string(data)}, " ")
}
