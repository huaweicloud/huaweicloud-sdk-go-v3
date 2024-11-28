package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VoiceLanguage **参数解释**： 声音语言。 **约束限制**： 不涉及。 **取值范围**： * UNKNOW：未知 * CN：中文 * EN：英文 * GER：德语 * fr：法语 * Kr：韩语 * por：葡萄牙语 * JPN：日语 * Ita：意大利语 * ESP：西班牙语 * DBH：东北话 * GT：港台 * GXH：广西话 * HBH：湖北话 * SXH：陕西话 * SCH：四川话 * YY：粤语 * Russian: 俄罗斯语 * Filipino: 菲律宾语 * Dutch: 荷兰语 * Indonesian: 印尼语 * Vietnamese: 越南语 * Arabic: 阿拉伯语 * Turkish: 土耳其语 * Malay: 马来语 * Thai: 泰语 * Finnish: 芬兰语
type VoiceLanguage struct {
	value string
}

type VoiceLanguageEnum struct {
	UNKNOW     VoiceLanguage
	CN         VoiceLanguage
	EN         VoiceLanguage
	GER        VoiceLanguage
	FR         VoiceLanguage
	KR         VoiceLanguage
	POR        VoiceLanguage
	JPN        VoiceLanguage
	ITA        VoiceLanguage
	ESP        VoiceLanguage
	DBH        VoiceLanguage
	GT         VoiceLanguage
	GXH        VoiceLanguage
	HBH        VoiceLanguage
	SXH        VoiceLanguage
	SCH        VoiceLanguage
	YY         VoiceLanguage
	RUSSIAN    VoiceLanguage
	FILIPINO   VoiceLanguage
	DUTCH      VoiceLanguage
	INDONESIAN VoiceLanguage
	VIETNAMESE VoiceLanguage
	ARABIC     VoiceLanguage
	TURKISH    VoiceLanguage
	MALAY      VoiceLanguage
	THAI       VoiceLanguage
	FINNISH    VoiceLanguage
}

func GetVoiceLanguageEnum() VoiceLanguageEnum {
	return VoiceLanguageEnum{
		UNKNOW: VoiceLanguage{
			value: "UNKNOW",
		},
		CN: VoiceLanguage{
			value: "CN",
		},
		EN: VoiceLanguage{
			value: "EN",
		},
		GER: VoiceLanguage{
			value: "GER",
		},
		FR: VoiceLanguage{
			value: "fr",
		},
		KR: VoiceLanguage{
			value: "Kr",
		},
		POR: VoiceLanguage{
			value: "por",
		},
		JPN: VoiceLanguage{
			value: "JPN",
		},
		ITA: VoiceLanguage{
			value: "Ita",
		},
		ESP: VoiceLanguage{
			value: "ESP",
		},
		DBH: VoiceLanguage{
			value: "DBH",
		},
		GT: VoiceLanguage{
			value: "GT",
		},
		GXH: VoiceLanguage{
			value: "GXH",
		},
		HBH: VoiceLanguage{
			value: "HBH",
		},
		SXH: VoiceLanguage{
			value: "SXH",
		},
		SCH: VoiceLanguage{
			value: "SCH",
		},
		YY: VoiceLanguage{
			value: "YY",
		},
		RUSSIAN: VoiceLanguage{
			value: "Russian",
		},
		FILIPINO: VoiceLanguage{
			value: "Filipino",
		},
		DUTCH: VoiceLanguage{
			value: "Dutch",
		},
		INDONESIAN: VoiceLanguage{
			value: "Indonesian",
		},
		VIETNAMESE: VoiceLanguage{
			value: "Vietnamese",
		},
		ARABIC: VoiceLanguage{
			value: "Arabic",
		},
		TURKISH: VoiceLanguage{
			value: "Turkish",
		},
		MALAY: VoiceLanguage{
			value: "Malay",
		},
		THAI: VoiceLanguage{
			value: "Thai",
		},
		FINNISH: VoiceLanguage{
			value: "Finnish",
		},
	}
}

func (c VoiceLanguage) Value() string {
	return c.value
}

func (c VoiceLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceLanguage) UnmarshalJSON(b []byte) error {
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
