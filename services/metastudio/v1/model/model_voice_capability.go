package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VoiceCapability 音色资产支持的能力集。 > 音色能力集只允许查询，不允许设置
type VoiceCapability struct {

	// **参数解释**： 该声音是否支持英文音标。 **约束限制**： 不涉及 **取值范围**： * true: 支持英文音标 * false: 不支持英文音标
	IsSupportPhonemeEn *bool `json:"is_support_phoneme_en,omitempty"`

	// **参数解释**： 该声音是否支持中文多音字。 **约束限制**： 不涉及 **取值范围**： * true: 支持中文多音字 * false: 不支持中文多音字
	IsSupportPhoneme *bool `json:"is_support_phoneme,omitempty"`

	// **参数解释**： 该声音是否支持停顿。 **约束限制**： 不涉及 **取值范围**： * true: 支持停顿 * false: 不支持停顿
	IsSupportBreakTime *bool `json:"is_support_break_time,omitempty"`

	// **参数解释**： 该声音是否支持韵律。 **约束限制**： 不涉及 **取值范围**： * true: 支持韵律 * false: 不支持韵律
	IsSupportBreakStrength *bool `json:"is_support_break_strength,omitempty"`

	// **参数解释**： 该声音是否支持全局语速。 **约束限制**： 不涉及 **取值范围**： * true: 支持全局语速 * false: 不支持全局语速
	IsSupportSpeed *bool `json:"is_support_speed,omitempty"`

	// **参数解释**： 该声音是否支持局部语速。 **约束限制**： 不涉及 **取值范围**： * true: 支持局部语速 * false: 不支持局部语速
	IsSupportProsody *bool `json:"is_support_prosody,omitempty"`

	// **参数解释**： 该声音是否支持SSML的say-as标签。 **约束限制**： 不涉及 **取值范围**： * true: 支持SSML的say-as标签 * false: 不支持SSML的say-as标签
	IsSupportSsmlSayAs *bool `json:"is_support_ssml_say_as,omitempty"`

	// **参数解释**： 该声音是否支持SSML的sub标签。 **约束限制**： 不涉及 **取值范围**： * true: 支持SSML的sub标签 * false: 不支持SSML的sub标签
	IsSupportSsmlSub *bool `json:"is_support_ssml_sub,omitempty"`

	// **参数解释**： 该声音是否支持连读。 **约束限制**： 不涉及 **取值范围**： * true: 支持连读 * false: 不支持连读
	IsSupportWord *bool `json:"is_support_word,omitempty"`

	// 是否支持缓存。
	IsSupportVoiceCache *bool `json:"is_support_voice_cache,omitempty"`

	// **参数解释**： 合成率。 **约束限制**： 不涉及 **取值范围**： * 0-50
	ConversionRate *float32 `json:"conversion_rate,omitempty"`

	// **参数解释**： 英语的合成率。 **约束限制**： 不涉及 **取值范围**： * 0-50
	ConversionRateEn *float32 `json:"conversion_rate_en,omitempty"`
}

func (o VoiceCapability) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceCapability struct{}"
	}

	return strings.Join([]string{"VoiceCapability", string(data)}, " ")
}
