package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VoiceCapability 音色资产支持的能力集。 > 音色能力集只允许查询，不允许设置
type VoiceCapability struct {

	// 支持英文音标。
	IsSupportPhonemeEn *bool `json:"is_support_phoneme_en,omitempty"`

	// 是否支持多音字。
	IsSupportPhoneme *bool `json:"is_support_phoneme,omitempty"`

	// 是否支持停顿。
	IsSupportBreakTime *bool `json:"is_support_break_time,omitempty"`

	// 是否支持韵律。
	IsSupportBreakStrength *bool `json:"is_support_break_strength,omitempty"`

	// 是否支持全局语速。
	IsSupportSpeed *bool `json:"is_support_speed,omitempty"`

	// 是否支持局部语速。
	IsSupportProsody *bool `json:"is_support_prosody,omitempty"`

	// 是否支持SSML的say-as标签。
	IsSupportSsmlSayAs *bool `json:"is_support_ssml_say_as,omitempty"`

	// 是否支持SSML的sub标签。
	IsSupportSsmlSub *bool `json:"is_support_ssml_sub,omitempty"`

	// 是否支持连读。
	IsSupportWord *bool `json:"is_support_word,omitempty"`
}

func (o VoiceCapability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceCapability struct{}"
	}

	return strings.Join([]string{"VoiceCapability", string(data)}, " ")
}
