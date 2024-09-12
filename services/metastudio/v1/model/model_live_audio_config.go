package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveAudioConfig 直播音频配置
type LiveAudioConfig struct {

	// **参数解释**： 插入音频资产的资产id，外部资产信息无需填写。若填写，可以从资产库中查询。 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及。
	AssetId *string `json:"asset_id,omitempty"`

	// **参数解释**： 音频URL。 **约束限制**： 仅支持MP3格式，大小<100MB。输出会自动转化为单声道16KHZ采样。 **取值范围**： 字符长度0-2048位。 **默认取值**： 不涉及。
	AudioUrl *string `json:"audio_url,omitempty"`

	// **参数解释**： 音频对应的字幕文件URL。 **约束限制**： 仅SRT格式，大小<1MB。 **取值范围**： 字符长度0-2048位。 **默认取值**： 不涉及。
	SubtitleUrl *string `json:"subtitle_url,omitempty"`
}

func (o LiveAudioConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveAudioConfig struct{}"
	}

	return strings.Join([]string{"LiveAudioConfig", string(data)}, " ")
}
