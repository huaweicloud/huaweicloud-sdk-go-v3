package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// XimalayaVoiceAssetMeta 喜马拉雅TTS音色元数据。此参数仅内部使用，不对外开放。
type XimalayaVoiceAssetMeta struct {

	// 音色适用领域。
	Domain string `json:"domain"`

	// 音色名称。
	SpeakerName string `json:"speaker_name"`

	// 音色变声。
	SpeakerVariant string `json:"speaker_variant"`
}

func (o XimalayaVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "XimalayaVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"XimalayaVoiceAssetMeta", string(data)}, " ")
}
