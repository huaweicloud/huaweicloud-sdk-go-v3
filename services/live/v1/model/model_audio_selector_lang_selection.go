package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AudioSelectorLangSelection 音频选择器语言设置
type AudioSelectorLangSelection struct {

	// 语言简称，输入2或3个小写字母的语言代码。
	LanguageCode string `json:"language_code"`

	// 语言输出策略。  取值如下： - LOOSE：宽松匹配，示例“eng”会优先匹配源流中语言为English的音轨，如果匹配不到，则选择PID最小的音轨。 - STRICT：严格匹配，示例“eng”会严格匹配源流中语言为English的音轨，如果匹配不到，媒体直播服务会自动补齐一个静音分片，当终端使用此音频选择器播放视频时，会静音播放。
	LanguageSelectionPolicy *AudioSelectorLangSelectionLanguageSelectionPolicy `json:"language_selection_policy,omitempty"`
}

func (o AudioSelectorLangSelection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioSelectorLangSelection struct{}"
	}

	return strings.Join([]string{"AudioSelectorLangSelection", string(data)}, " ")
}

type AudioSelectorLangSelectionLanguageSelectionPolicy struct {
	value string
}

type AudioSelectorLangSelectionLanguageSelectionPolicyEnum struct {
	LOOSE  AudioSelectorLangSelectionLanguageSelectionPolicy
	STRICT AudioSelectorLangSelectionLanguageSelectionPolicy
}

func GetAudioSelectorLangSelectionLanguageSelectionPolicyEnum() AudioSelectorLangSelectionLanguageSelectionPolicyEnum {
	return AudioSelectorLangSelectionLanguageSelectionPolicyEnum{
		LOOSE: AudioSelectorLangSelectionLanguageSelectionPolicy{
			value: "LOOSE",
		},
		STRICT: AudioSelectorLangSelectionLanguageSelectionPolicy{
			value: "STRICT",
		},
	}
}

func (c AudioSelectorLangSelectionLanguageSelectionPolicy) Value() string {
	return c.value
}

func (c AudioSelectorLangSelectionLanguageSelectionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioSelectorLangSelectionLanguageSelectionPolicy) UnmarshalJSON(b []byte) error {
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
