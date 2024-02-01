package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MobvoiVoiceAssetMeta 出门问问 TTS音色元数据。
type MobvoiVoiceAssetMeta struct {

	// 合成音频指定发音人
	Speaker string `json:"speaker"`
}

func (o MobvoiVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MobvoiVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"MobvoiVoiceAssetMeta", string(data)}, " ")
}
