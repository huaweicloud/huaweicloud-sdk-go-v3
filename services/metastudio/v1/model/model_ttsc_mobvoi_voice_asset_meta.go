package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtscMobvoiVoiceAssetMeta 出门问问TTS音色元数据。
type TtscMobvoiVoiceAssetMeta struct {

	// 合成音频指定发音人
	Speaker string `json:"speaker"`
}

func (o TtscMobvoiVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtscMobvoiVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"TtscMobvoiVoiceAssetMeta", string(data)}, " ")
}
