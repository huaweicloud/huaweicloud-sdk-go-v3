package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VoiceConfig 语音配置参数。
type VoiceConfig struct {

	// **参数解释**： 音色资产ID，可以从资产库中查询。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	VoiceAssetId string `json:"voice_asset_id"`

	// **参数解释**： 语速。50表示0.5倍语速，100表示正常语速，200表示2倍语速。 当取值为“100”时，表示一个成年人的正常语速，约为250字/分钟。  **约束限制**： 不涉及。
	Speed *int32 `json:"speed,omitempty"`

	// **参数解释**： 音高。 **约束限制**： 不涉及。
	Pitch *int32 `json:"pitch,omitempty"`

	// **参数解释**： 音量。 **约束限制**： 不涉及。
	Volume *int32 `json:"volume,omitempty"`
}

func (o VoiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceConfig struct{}"
	}

	return strings.Join([]string{"VoiceConfig", string(data)}, " ")
}
