package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtscExternalVoiceAssetMeta 第三方TTS音色元数据。
type TtscExternalVoiceAssetMeta struct {

	// 第三方TTS供应商类型。 * XIMALAYA：喜马拉雅TTS * HUAWEI_EI：EI TTS * MOBVOI：出门问问TTS
	Provider string `json:"provider"`

	MobvoiVoiceMeta *TtscMobvoiVoiceAssetMeta `json:"mobvoi_voice_meta,omitempty"`

	AudioxVoiceMeta *TtscAudioxVoiceAssetMeta `json:"audiox_voice_meta,omitempty"`
}

func (o TtscExternalVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtscExternalVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"TtscExternalVoiceAssetMeta", string(data)}, " ")
}
