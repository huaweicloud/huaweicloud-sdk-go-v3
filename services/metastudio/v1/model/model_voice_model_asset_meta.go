package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 音色模型元数据。
type VoiceModelAssetMeta struct {

	// 音色资产类型。 * COMMON：通用情感模型 * CLONE：语音克隆模型
	ModelType *VoiceModelAssetMetaModelType `json:"model_type,omitempty"`

	// 音色性别。默认UNKNOW。 * UNKNOW： 中性音色 * MALE： 男性音色 * FEMALE： 女性音色
	Sex *VoiceModelAssetMetaSex `json:"sex,omitempty"`

	// 音色语言。默认UNKNOW。 * UNKNOW： 未知 * CN： 中文 * EN： 英文
	Language *VoiceModelAssetMetaLanguage `json:"language,omitempty"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
