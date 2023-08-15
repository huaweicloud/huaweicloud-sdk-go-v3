package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HuaweiEiVoiceAssetMeta 华为云EI TTS音色元数据。此参数仅内部使用，不对外开放。
type HuaweiEiVoiceAssetMeta struct {

	// 音色属性。
	Property string `json:"property"`

	// 是否是克隆音色，默认是false。
	IsClonedVoice *bool `json:"is_cloned_voice,omitempty"`

	// 音色克隆时的训练任务ID。当is_cloned_voice=true时需要填写。
	TrainingJobId *string `json:"training_job_id,omitempty"`
}

func (o HuaweiEiVoiceAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HuaweiEiVoiceAssetMeta struct{}"
	}

	return strings.Join([]string{"HuaweiEiVoiceAssetMeta", string(data)}, " ")
}
