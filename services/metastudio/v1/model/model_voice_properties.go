package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VoiceProperties 声音配置信息
type VoiceProperties struct {

	// 任务标签。   * ECOMMERCE: 电商   * NEWS: 新闻   * MARKETING: 营销
	JobTag VoicePropertiesJobTag `json:"job_tag"`

	// 语音性别,是男性声音还是女性声音。 * FEMALE: 女性 * MALE: 男性
	Sex VoicePropertiesSex `json:"sex"`

	// 训练语言,当前仅支持中文。 * CN: 中文 * EN: 英文
	Language VoicePropertiesLanguage `json:"language"`
}

func (o VoiceProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceProperties struct{}"
	}

	return strings.Join([]string{"VoiceProperties", string(data)}, " ")
}

type VoicePropertiesJobTag struct {
	value string
}

type VoicePropertiesJobTagEnum struct {
	ECOMMERCE VoicePropertiesJobTag
	NEWS      VoicePropertiesJobTag
	MARKETING VoicePropertiesJobTag
}

func GetVoicePropertiesJobTagEnum() VoicePropertiesJobTagEnum {
	return VoicePropertiesJobTagEnum{
		ECOMMERCE: VoicePropertiesJobTag{
			value: "ECOMMERCE",
		},
		NEWS: VoicePropertiesJobTag{
			value: "NEWS",
		},
		MARKETING: VoicePropertiesJobTag{
			value: "MARKETING",
		},
	}
}

func (c VoicePropertiesJobTag) Value() string {
	return c.value
}

func (c VoicePropertiesJobTag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoicePropertiesJobTag) UnmarshalJSON(b []byte) error {
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

type VoicePropertiesSex struct {
	value string
}

type VoicePropertiesSexEnum struct {
	FEMALE VoicePropertiesSex
	MALE   VoicePropertiesSex
}

func GetVoicePropertiesSexEnum() VoicePropertiesSexEnum {
	return VoicePropertiesSexEnum{
		FEMALE: VoicePropertiesSex{
			value: "FEMALE",
		},
		MALE: VoicePropertiesSex{
			value: "MALE",
		},
	}
}

func (c VoicePropertiesSex) Value() string {
	return c.value
}

func (c VoicePropertiesSex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoicePropertiesSex) UnmarshalJSON(b []byte) error {
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

type VoicePropertiesLanguage struct {
	value string
}

type VoicePropertiesLanguageEnum struct {
	CN VoicePropertiesLanguage
	EN VoicePropertiesLanguage
}

func GetVoicePropertiesLanguageEnum() VoicePropertiesLanguageEnum {
	return VoicePropertiesLanguageEnum{
		CN: VoicePropertiesLanguage{
			value: "CN",
		},
		EN: VoicePropertiesLanguage{
			value: "EN",
		},
	}
}

func (c VoicePropertiesLanguage) Value() string {
	return c.value
}

func (c VoicePropertiesLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoicePropertiesLanguage) UnmarshalJSON(b []byte) error {
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
