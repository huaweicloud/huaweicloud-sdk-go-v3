package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VoiceModelAssetMeta 音色模型元数据。
type VoiceModelAssetMeta struct {

	// 音色资产类型。 * COMMON：通用情感模型 * CLONE：语音克隆模型
	ModelType *VoiceModelAssetMetaModelType `json:"model_type,omitempty"`

	// 音色性别。 * UNKNOW：中性音色 * MALE：男性音色 * FEMALE：女性音色  默认UNKNOW。
	Sex *VoiceModelAssetMetaSex `json:"sex,omitempty"`

	// 音色语言。 * UNKNOW：未知 * CN：中文 * EN：英文  默认UNKNOW。
	Language *VoiceModelAssetMetaLanguage `json:"language,omitempty"`

	// 自研TTS运行模式，包括CPU模式和GPU模式。此参数仅内部使用，不对外开放。 * UNKNOW：未知 * TTS_V1：V1版本TTS，运行在CPU上 * TTS_V2：V2版本TTS，运行在GPU上  默认UNKNOW。
	TtsMode *VoiceModelAssetMetaTtsMode `json:"tts_mode,omitempty"`

	ExternalVoiceMeta *ExternalVoiceAssetMeta `json:"external_voice_meta,omitempty"`
}

func (o VoiceModelAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceModelAssetMeta struct{}"
	}

	return strings.Join([]string{"VoiceModelAssetMeta", string(data)}, " ")
}

type VoiceModelAssetMetaModelType struct {
	value string
}

type VoiceModelAssetMetaModelTypeEnum struct {
	COMMON VoiceModelAssetMetaModelType
	CLONE  VoiceModelAssetMetaModelType
}

func GetVoiceModelAssetMetaModelTypeEnum() VoiceModelAssetMetaModelTypeEnum {
	return VoiceModelAssetMetaModelTypeEnum{
		COMMON: VoiceModelAssetMetaModelType{
			value: "COMMON",
		},
		CLONE: VoiceModelAssetMetaModelType{
			value: "CLONE",
		},
	}
}

func (c VoiceModelAssetMetaModelType) Value() string {
	return c.value
}

func (c VoiceModelAssetMetaModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceModelAssetMetaModelType) UnmarshalJSON(b []byte) error {
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

type VoiceModelAssetMetaSex struct {
	value string
}

type VoiceModelAssetMetaSexEnum struct {
	UNKNOW VoiceModelAssetMetaSex
	MALE   VoiceModelAssetMetaSex
	FEMALE VoiceModelAssetMetaSex
}

func GetVoiceModelAssetMetaSexEnum() VoiceModelAssetMetaSexEnum {
	return VoiceModelAssetMetaSexEnum{
		UNKNOW: VoiceModelAssetMetaSex{
			value: "UNKNOW",
		},
		MALE: VoiceModelAssetMetaSex{
			value: "MALE",
		},
		FEMALE: VoiceModelAssetMetaSex{
			value: "FEMALE",
		},
	}
}

func (c VoiceModelAssetMetaSex) Value() string {
	return c.value
}

func (c VoiceModelAssetMetaSex) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceModelAssetMetaSex) UnmarshalJSON(b []byte) error {
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

type VoiceModelAssetMetaLanguage struct {
	value string
}

type VoiceModelAssetMetaLanguageEnum struct {
	UNKNOW VoiceModelAssetMetaLanguage
	CN     VoiceModelAssetMetaLanguage
	EN     VoiceModelAssetMetaLanguage
}

func GetVoiceModelAssetMetaLanguageEnum() VoiceModelAssetMetaLanguageEnum {
	return VoiceModelAssetMetaLanguageEnum{
		UNKNOW: VoiceModelAssetMetaLanguage{
			value: "UNKNOW",
		},
		CN: VoiceModelAssetMetaLanguage{
			value: "CN",
		},
		EN: VoiceModelAssetMetaLanguage{
			value: "EN",
		},
	}
}

func (c VoiceModelAssetMetaLanguage) Value() string {
	return c.value
}

func (c VoiceModelAssetMetaLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceModelAssetMetaLanguage) UnmarshalJSON(b []byte) error {
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

type VoiceModelAssetMetaTtsMode struct {
	value string
}

type VoiceModelAssetMetaTtsModeEnum struct {
	UNKNOW VoiceModelAssetMetaTtsMode
	TTS_V1 VoiceModelAssetMetaTtsMode
	TTS_V2 VoiceModelAssetMetaTtsMode
}

func GetVoiceModelAssetMetaTtsModeEnum() VoiceModelAssetMetaTtsModeEnum {
	return VoiceModelAssetMetaTtsModeEnum{
		UNKNOW: VoiceModelAssetMetaTtsMode{
			value: "UNKNOW",
		},
		TTS_V1: VoiceModelAssetMetaTtsMode{
			value: "TTS_V1",
		},
		TTS_V2: VoiceModelAssetMetaTtsMode{
			value: "TTS_V2",
		},
	}
}

func (c VoiceModelAssetMetaTtsMode) Value() string {
	return c.value
}

func (c VoiceModelAssetMetaTtsMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceModelAssetMetaTtsMode) UnmarshalJSON(b []byte) error {
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
