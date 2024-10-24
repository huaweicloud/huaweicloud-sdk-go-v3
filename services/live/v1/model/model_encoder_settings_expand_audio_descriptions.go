package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EncoderSettingsExpandAudioDescriptions struct {

	// 音频输出配置的名称。仅支持大小写字母，数字，中划线（-），下划线（_）。  同一个频道不同的音频输出配置名称，不允许重复。
	Name string `json:"name"`

	// 音频选择器名称
	AudioSelectorName string `json:"audio_selector_name"`

	// 语言代码控制。这里的设置不会修改音频实际的语言，只会修改音频对外展示的语言。  包含如下选项： - FOLLOW_INPUT：如果所选音频选择器对应的输出音频有语言，则与其保持一致，否则会以这里配置的语言代码和流名称进行兜底。推荐当前选项，为默认值。 - USE_CONFIGURED：用户根据实际情况自定义输出音频的语言和流名称。
	LanguageCodeControl *EncoderSettingsExpandAudioDescriptionsLanguageCodeControl `json:"language_code_control,omitempty"`

	// 语言代码，输入2或3个小写字母。
	LanguageCode *string `json:"language_code,omitempty"`

	// 流名称
	StreamName *string `json:"stream_name,omitempty"`
}

func (o EncoderSettingsExpandAudioDescriptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncoderSettingsExpandAudioDescriptions struct{}"
	}

	return strings.Join([]string{"EncoderSettingsExpandAudioDescriptions", string(data)}, " ")
}

type EncoderSettingsExpandAudioDescriptionsLanguageCodeControl struct {
	value string
}

type EncoderSettingsExpandAudioDescriptionsLanguageCodeControlEnum struct {
	FOLLOW_INPUT   EncoderSettingsExpandAudioDescriptionsLanguageCodeControl
	USE_CONFIGURED EncoderSettingsExpandAudioDescriptionsLanguageCodeControl
}

func GetEncoderSettingsExpandAudioDescriptionsLanguageCodeControlEnum() EncoderSettingsExpandAudioDescriptionsLanguageCodeControlEnum {
	return EncoderSettingsExpandAudioDescriptionsLanguageCodeControlEnum{
		FOLLOW_INPUT: EncoderSettingsExpandAudioDescriptionsLanguageCodeControl{
			value: "FOLLOW_INPUT",
		},
		USE_CONFIGURED: EncoderSettingsExpandAudioDescriptionsLanguageCodeControl{
			value: "USE_CONFIGURED",
		},
	}
}

func (c EncoderSettingsExpandAudioDescriptionsLanguageCodeControl) Value() string {
	return c.value
}

func (c EncoderSettingsExpandAudioDescriptionsLanguageCodeControl) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EncoderSettingsExpandAudioDescriptionsLanguageCodeControl) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
