package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtscAudioxVoiceAssetMeta 逻辑智能TTS音色元数据。
type TtscAudioxVoiceAssetMeta struct {

	// 合成音频指定发音人
	Speaker string `json:"speaker"`
}

func (o TtscAudioxVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtscAudioxVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"TtscAudioxVoiceAssetMeta", string(data)}, " ")
}
