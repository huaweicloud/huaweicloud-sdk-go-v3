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

	// 音色语言。 * UNKNOW：未知 * CN：中文 * EN：英文 * GER：德语 * fr：法语 * Kr：韩语 * por：葡萄牙语 * JPN：日语 * Ita：意大利语 * ESP：西班牙语 * DBH：东北话 * GT：港台 * GXH：广西话 * HBH：湖北话 * SXH：陕西话 * SCH：四川话 * YY：粤语 * Russian: 俄罗斯语 * Filipino: 菲律宾语 * Dutch: 荷兰语 * Indonesian: 印尼语 * Vietnamese: 越南语 * Arabic: 阿拉伯语 * Turkish: 土耳其语 * Malay: 马来语 * Thai: 泰语 * Finnish: 芬兰语  默认UNKNOW。
	Language *VoiceModelAssetMetaLanguage `json:"language,omitempty"`

	// 语速缩放比例
	SpeedRatio *float32 `json:"speed_ratio,omitempty"`

	// 音量缩放比例
	VolumeRatio *float32 `json:"volume_ratio,omitempty"`
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
	UNKNOW     VoiceModelAssetMetaLanguage
	CN         VoiceModelAssetMetaLanguage
	EN         VoiceModelAssetMetaLanguage
	GER        VoiceModelAssetMetaLanguage
	FR         VoiceModelAssetMetaLanguage
	KR         VoiceModelAssetMetaLanguage
	POR        VoiceModelAssetMetaLanguage
	JPN        VoiceModelAssetMetaLanguage
	ITA        VoiceModelAssetMetaLanguage
	ESP        VoiceModelAssetMetaLanguage
	DBH        VoiceModelAssetMetaLanguage
	GT         VoiceModelAssetMetaLanguage
	GXH        VoiceModelAssetMetaLanguage
	HBH        VoiceModelAssetMetaLanguage
	SXH        VoiceModelAssetMetaLanguage
	SCH        VoiceModelAssetMetaLanguage
	YY         VoiceModelAssetMetaLanguage
	RUSSIAN    VoiceModelAssetMetaLanguage
	FILIPINO   VoiceModelAssetMetaLanguage
	DUTCH      VoiceModelAssetMetaLanguage
	INDONESIAN VoiceModelAssetMetaLanguage
	VIETNAMESE VoiceModelAssetMetaLanguage
	ARABIC     VoiceModelAssetMetaLanguage
	TURKISH    VoiceModelAssetMetaLanguage
	MALAY      VoiceModelAssetMetaLanguage
	THAI       VoiceModelAssetMetaLanguage
	FINNISH    VoiceModelAssetMetaLanguage
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
		GER: VoiceModelAssetMetaLanguage{
			value: "GER",
		},
		FR: VoiceModelAssetMetaLanguage{
			value: "fr",
		},
		KR: VoiceModelAssetMetaLanguage{
			value: "Kr",
		},
		POR: VoiceModelAssetMetaLanguage{
			value: "por",
		},
		JPN: VoiceModelAssetMetaLanguage{
			value: "JPN",
		},
		ITA: VoiceModelAssetMetaLanguage{
			value: "Ita",
		},
		ESP: VoiceModelAssetMetaLanguage{
			value: "ESP",
		},
		DBH: VoiceModelAssetMetaLanguage{
			value: "DBH",
		},
		GT: VoiceModelAssetMetaLanguage{
			value: "GT",
		},
		GXH: VoiceModelAssetMetaLanguage{
			value: "GXH",
		},
		HBH: VoiceModelAssetMetaLanguage{
			value: "HBH",
		},
		SXH: VoiceModelAssetMetaLanguage{
			value: "SXH",
		},
		SCH: VoiceModelAssetMetaLanguage{
			value: "SCH",
		},
		YY: VoiceModelAssetMetaLanguage{
			value: "YY",
		},
		RUSSIAN: VoiceModelAssetMetaLanguage{
			value: "Russian",
		},
		FILIPINO: VoiceModelAssetMetaLanguage{
			value: "Filipino",
		},
		DUTCH: VoiceModelAssetMetaLanguage{
			value: "Dutch",
		},
		INDONESIAN: VoiceModelAssetMetaLanguage{
			value: "Indonesian",
		},
		VIETNAMESE: VoiceModelAssetMetaLanguage{
			value: "Vietnamese",
		},
		ARABIC: VoiceModelAssetMetaLanguage{
			value: "Arabic",
		},
		TURKISH: VoiceModelAssetMetaLanguage{
			value: "Turkish",
		},
		MALAY: VoiceModelAssetMetaLanguage{
			value: "Malay",
		},
		THAI: VoiceModelAssetMetaLanguage{
			value: "Thai",
		},
		FINNISH: VoiceModelAssetMetaLanguage{
			value: "Finnish",
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
