package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioAssetMeta 音频元数据，自动提取获得。
type AudioAssetMeta struct {

	// **参数解释**： 时长,单位秒。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**： 音频编码格式。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： 字符长度0-32位。 **默认取值**： 不涉及
	AudioCodec *string `json:"audio_codec,omitempty"`

	// **参数解释**： 音频平均码率,单位kbps。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	AudioBitRate *int32 `json:"audio_bit_rate,omitempty"`

	// **参数解释**： 音频声道数。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	AudioChannels *int32 `json:"audio_channels,omitempty"`

	// **参数解释**： 采样率,HZ。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	Sample *int32 `json:"sample,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o AudioAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioAssetMeta struct{}"
	}

	return strings.Join([]string{"AudioAssetMeta", string(data)}, " ")
}
