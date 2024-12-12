package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChatVoiceConfig 语音配置参数
type ChatVoiceConfig struct {

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

	// 语言类型。默认值CN。 * CN：简体中文。 * EN：英语。
	Language *ChatVoiceConfigLanguage `json:"language,omitempty"`
}

func (o ChatVoiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatVoiceConfig struct{}"
	}

	return strings.Join([]string{"ChatVoiceConfig", string(data)}, " ")
}

type ChatVoiceConfigLanguage struct {
	value string
}

type ChatVoiceConfigLanguageEnum struct {
	CN ChatVoiceConfigLanguage
	EN ChatVoiceConfigLanguage
}

func GetChatVoiceConfigLanguageEnum() ChatVoiceConfigLanguageEnum {
	return ChatVoiceConfigLanguageEnum{
		CN: ChatVoiceConfigLanguage{
			value: "CN",
		},
		EN: ChatVoiceConfigLanguage{
			value: "EN",
		},
	}
}

func (c ChatVoiceConfigLanguage) Value() string {
	return c.value
}

func (c ChatVoiceConfigLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChatVoiceConfigLanguage) UnmarshalJSON(b []byte) error {
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
