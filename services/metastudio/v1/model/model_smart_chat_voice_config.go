package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmartChatVoiceConfig 语音配置参数
type SmartChatVoiceConfig struct {

	// 语音合成特征字符串
	VoiceAssetId *string `json:"voice_asset_id,omitempty"`

	// 语速。默认值100，最小值50，最大值200。 > 当取值为“100”时，表示一个成年人正常的语速，约为250字/分钟。
	Speed *int32 `json:"speed,omitempty"`

	// 音高。默认值100，最小值50，最大值200。
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。默认值140，最小值90，最大值240。
	Volume *int32 `json:"volume,omitempty"`

	// 第三方TTS供应商类型。 * XIMALAYA：喜马拉雅TTS * HUAWEI_EI：EI TTS * MOBVOI：出门问问TTS
	Provider *string `json:"provider,omitempty"`

	// 语言类型。默认值CN。 * CN：中文。 * EN：英文。
	Language *SmartChatVoiceConfigLanguage `json:"language,omitempty"`

	// 语言描述。
	LanguageDesc *string `json:"language_desc,omitempty"`
}

func (o SmartChatVoiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatVoiceConfig struct{}"
	}

	return strings.Join([]string{"SmartChatVoiceConfig", string(data)}, " ")
}

type SmartChatVoiceConfigLanguage struct {
	value string
}

type SmartChatVoiceConfigLanguageEnum struct {
	CN SmartChatVoiceConfigLanguage
	EN SmartChatVoiceConfigLanguage
}

func GetSmartChatVoiceConfigLanguageEnum() SmartChatVoiceConfigLanguageEnum {
	return SmartChatVoiceConfigLanguageEnum{
		CN: SmartChatVoiceConfigLanguage{
			value: "CN",
		},
		EN: SmartChatVoiceConfigLanguage{
			value: "EN",
		},
	}
}

func (c SmartChatVoiceConfigLanguage) Value() string {
	return c.value
}

func (c SmartChatVoiceConfigLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartChatVoiceConfigLanguage) UnmarshalJSON(b []byte) error {
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
