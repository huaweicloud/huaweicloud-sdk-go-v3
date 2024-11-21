package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Language **参数解释**： 声音语言。 **约束限制**： 不涉及。 **取值范围**： * UNKNOW：未知 * CN：中文 * EN：英文 * GER：德语 * fr：法语 * Kr：韩语 * por：葡萄牙语 * JPN：日语 * Ita：意大利语 * ESP：西班牙语 * DBH：东北话 * GT：港台 * GXH：广西话 * HBH：湖北话 * SXH：陕西话 * SCH：四川话 * YY：粤语 * Russian: 俄罗斯语 * Filipino: 菲律宾语 * Dutch: 荷兰语 * Indonesian: 印尼语 * Vietnamese: 越南语 * Arabic: 阿拉伯语 * Turkish: 土耳其语 * Malay: 马来语 * Thai: 泰语 * Finnish: 芬兰语
type Language struct {
	value string
}

type LanguageEnum struct {
	UNKNOW     Language
	CN         Language
	EN         Language
	GER        Language
	FR         Language
	KR         Language
	POR        Language
	JPN        Language
	ITA        Language
	ESP        Language
	DBH        Language
	GT         Language
	GXH        Language
	HBH        Language
	SXH        Language
	SCH        Language
	YY         Language
	RUSSIAN    Language
	FILIPINO   Language
	DUTCH      Language
	INDONESIAN Language
	VIETNAMESE Language
	ARABIC     Language
	TURKISH    Language
	MALAY      Language
	THAI       Language
	FINNISH    Language
}

func GetLanguageEnum() LanguageEnum {
	return LanguageEnum{
		UNKNOW: Language{
			value: "UNKNOW",
		},
		CN: Language{
			value: "CN",
		},
		EN: Language{
			value: "EN",
		},
		GER: Language{
			value: "GER",
		},
		FR: Language{
			value: "fr",
		},
		KR: Language{
			value: "Kr",
		},
		POR: Language{
			value: "por",
		},
		JPN: Language{
			value: "JPN",
		},
		ITA: Language{
			value: "Ita",
		},
		ESP: Language{
			value: "ESP",
		},
		DBH: Language{
			value: "DBH",
		},
		GT: Language{
			value: "GT",
		},
		GXH: Language{
			value: "GXH",
		},
		HBH: Language{
			value: "HBH",
		},
		SXH: Language{
			value: "SXH",
		},
		SCH: Language{
			value: "SCH",
		},
		YY: Language{
			value: "YY",
		},
		RUSSIAN: Language{
			value: "Russian",
		},
		FILIPINO: Language{
			value: "Filipino",
		},
		DUTCH: Language{
			value: "Dutch",
		},
		INDONESIAN: Language{
			value: "Indonesian",
		},
		VIETNAMESE: Language{
			value: "Vietnamese",
		},
		ARABIC: Language{
			value: "Arabic",
		},
		TURKISH: Language{
			value: "Turkish",
		},
		MALAY: Language{
			value: "Malay",
		},
		THAI: Language{
			value: "Thai",
		},
		FINNISH: Language{
			value: "Finnish",
		},
	}
}

func (c Language) Value() string {
	return c.value
}

func (c Language) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Language) UnmarshalJSON(b []byte) error {
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
