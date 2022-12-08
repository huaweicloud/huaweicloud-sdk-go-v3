package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type TextTranslationReq struct {

	// 待翻译文本，仅支持utf-8编码，长度不超过2000字符。
	Text string `json:"text"`

	// 翻译原语言，具体取值见支持的语言列表: 阿拉伯语 ar 德语 de 俄语 ru 法语 fr 韩语 ko 葡萄牙语 pt 日语 ja 泰语 th 土耳其语 tr 西班牙语 es 英语 en 越南语 vi 中文（简体） zh 中文（繁体） zh-tw 自动检测输入语种并翻译成目标语种，您需要指定目标语种。 auto
	From TextTranslationReqFrom `json:"from"`

	// 翻译目标语言，具体取值见支持的语言列表: 阿拉伯语 ar 德语 de 俄语 ru 法语 fr 韩语 ko 葡萄牙语 pt 日语 ja 泰语 th 土耳其语 tr 西班牙语 es 英语 en 越南语 vi 中文（简体） zh 中文（繁体） zh-tw
	To TextTranslationReqTo `json:"to"`

	// 默认为“common”，当前只有通用场景
	Scene *TextTranslationReqScene `json:"scene,omitempty"`
}

func (o TextTranslationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextTranslationReq struct{}"
	}

	return strings.Join([]string{"TextTranslationReq", string(data)}, " ")
}

type TextTranslationReqFrom struct {
	value string
}

type TextTranslationReqFromEnum struct {
	ZH    TextTranslationReqFrom
	EN    TextTranslationReqFrom
	JA    TextTranslationReqFrom
	RU    TextTranslationReqFrom
	ES    TextTranslationReqFrom
	DE    TextTranslationReqFrom
	AR    TextTranslationReqFrom
	PT    TextTranslationReqFrom
	TH    TextTranslationReqFrom
	TR    TextTranslationReqFrom
	KO    TextTranslationReqFrom
	FR    TextTranslationReqFrom
	VI    TextTranslationReqFrom
	HI    TextTranslationReqFrom
	KM    TextTranslationReqFrom
	MY    TextTranslationReqFrom
	TA    TextTranslationReqFrom
	FA    TextTranslationReqFrom
	UR    TextTranslationReqFrom
	BN    TextTranslationReqFrom
	MR    TextTranslationReqFrom
	GU    TextTranslationReqFrom
	PA    TextTranslationReqFrom
	TE    TextTranslationReqFrom
	HE    TextTranslationReqFrom
	CS    TextTranslationReqFrom
	SK    TextTranslationReqFrom
	RO    TextTranslationReqFrom
	SQ    TextTranslationReqFrom
	LV    TextTranslationReqFrom
	ET    TextTranslationReqFrom
	LT    TextTranslationReqFrom
	HR    TextTranslationReqFrom
	BS    TextTranslationReqFrom
	KA    TextTranslationReqFrom
	SL    TextTranslationReqFrom
	CA    TextTranslationReqFrom
	AF    TextTranslationReqFrom
	MG    TextTranslationReqFrom
	ID    TextTranslationReqFrom
	FIL   TextTranslationReqFrom
	SW    TextTranslationReqFrom
	HU    TextTranslationReqFrom
	SR    TextTranslationReqFrom
	MK    TextTranslationReqFrom
	UK    TextTranslationReqFrom
	BG    TextTranslationReqFrom
	MT    TextTranslationReqFrom
	EL    TextTranslationReqFrom
	IS    TextTranslationReqFrom
	GA    TextTranslationReqFrom
	CY    TextTranslationReqFrom
	HT    TextTranslationReqFrom
	NO    TextTranslationReqFrom
	SV    TextTranslationReqFrom
	FI    TextTranslationReqFrom
	DA    TextTranslationReqFrom
	ZH_TW TextTranslationReqFrom
	AUTO  TextTranslationReqFrom
}

func GetTextTranslationReqFromEnum() TextTranslationReqFromEnum {
	return TextTranslationReqFromEnum{
		ZH: TextTranslationReqFrom{
			value: "zh",
		},
		EN: TextTranslationReqFrom{
			value: "en",
		},
		JA: TextTranslationReqFrom{
			value: "ja",
		},
		RU: TextTranslationReqFrom{
			value: "ru",
		},
		ES: TextTranslationReqFrom{
			value: "es",
		},
		DE: TextTranslationReqFrom{
			value: "de",
		},
		AR: TextTranslationReqFrom{
			value: "ar",
		},
		PT: TextTranslationReqFrom{
			value: "pt",
		},
		TH: TextTranslationReqFrom{
			value: "th",
		},
		TR: TextTranslationReqFrom{
			value: "tr",
		},
		KO: TextTranslationReqFrom{
			value: "ko",
		},
		FR: TextTranslationReqFrom{
			value: "fr",
		},
		VI: TextTranslationReqFrom{
			value: "vi",
		},
		HI: TextTranslationReqFrom{
			value: "hi",
		},
		KM: TextTranslationReqFrom{
			value: "km",
		},
		MY: TextTranslationReqFrom{
			value: "my",
		},
		TA: TextTranslationReqFrom{
			value: "ta",
		},
		FA: TextTranslationReqFrom{
			value: "fa",
		},
		UR: TextTranslationReqFrom{
			value: "ur",
		},
		BN: TextTranslationReqFrom{
			value: "bn",
		},
		MR: TextTranslationReqFrom{
			value: "mr",
		},
		GU: TextTranslationReqFrom{
			value: "gu",
		},
		PA: TextTranslationReqFrom{
			value: "pa",
		},
		TE: TextTranslationReqFrom{
			value: "te",
		},
		HE: TextTranslationReqFrom{
			value: "he",
		},
		CS: TextTranslationReqFrom{
			value: "cs",
		},
		SK: TextTranslationReqFrom{
			value: "sk",
		},
		RO: TextTranslationReqFrom{
			value: "ro",
		},
		SQ: TextTranslationReqFrom{
			value: "sq",
		},
		LV: TextTranslationReqFrom{
			value: "lv",
		},
		ET: TextTranslationReqFrom{
			value: "et",
		},
		LT: TextTranslationReqFrom{
			value: "lt",
		},
		HR: TextTranslationReqFrom{
			value: "hr",
		},
		BS: TextTranslationReqFrom{
			value: "bs",
		},
		KA: TextTranslationReqFrom{
			value: "ka",
		},
		SL: TextTranslationReqFrom{
			value: "sl",
		},
		CA: TextTranslationReqFrom{
			value: "ca",
		},
		AF: TextTranslationReqFrom{
			value: "af",
		},
		MG: TextTranslationReqFrom{
			value: "mg",
		},
		ID: TextTranslationReqFrom{
			value: "id",
		},
		FIL: TextTranslationReqFrom{
			value: "fil",
		},
		SW: TextTranslationReqFrom{
			value: "sw",
		},
		HU: TextTranslationReqFrom{
			value: "hu",
		},
		SR: TextTranslationReqFrom{
			value: "sr",
		},
		MK: TextTranslationReqFrom{
			value: "mk",
		},
		UK: TextTranslationReqFrom{
			value: "uk",
		},
		BG: TextTranslationReqFrom{
			value: "bg",
		},
		MT: TextTranslationReqFrom{
			value: "mt",
		},
		EL: TextTranslationReqFrom{
			value: "el",
		},
		IS: TextTranslationReqFrom{
			value: "is",
		},
		GA: TextTranslationReqFrom{
			value: "ga",
		},
		CY: TextTranslationReqFrom{
			value: "cy",
		},
		HT: TextTranslationReqFrom{
			value: "ht",
		},
		NO: TextTranslationReqFrom{
			value: "no",
		},
		SV: TextTranslationReqFrom{
			value: "sv",
		},
		FI: TextTranslationReqFrom{
			value: "fi",
		},
		DA: TextTranslationReqFrom{
			value: "da",
		},
		ZH_TW: TextTranslationReqFrom{
			value: "zh-tw",
		},
		AUTO: TextTranslationReqFrom{
			value: "auto",
		},
	}
}

func (c TextTranslationReqFrom) Value() string {
	return c.value
}

func (c TextTranslationReqFrom) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextTranslationReqFrom) UnmarshalJSON(b []byte) error {
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

type TextTranslationReqTo struct {
	value string
}

type TextTranslationReqToEnum struct {
	ZH    TextTranslationReqTo
	EN    TextTranslationReqTo
	JA    TextTranslationReqTo
	RU    TextTranslationReqTo
	ES    TextTranslationReqTo
	DE    TextTranslationReqTo
	AR    TextTranslationReqTo
	PT    TextTranslationReqTo
	TH    TextTranslationReqTo
	TR    TextTranslationReqTo
	KO    TextTranslationReqTo
	FR    TextTranslationReqTo
	VI    TextTranslationReqTo
	HI    TextTranslationReqTo
	KM    TextTranslationReqTo
	MY    TextTranslationReqTo
	TA    TextTranslationReqTo
	FA    TextTranslationReqTo
	UR    TextTranslationReqTo
	BN    TextTranslationReqTo
	MR    TextTranslationReqTo
	GU    TextTranslationReqTo
	PA    TextTranslationReqTo
	TE    TextTranslationReqTo
	HE    TextTranslationReqTo
	CS    TextTranslationReqTo
	SK    TextTranslationReqTo
	RO    TextTranslationReqTo
	SQ    TextTranslationReqTo
	LV    TextTranslationReqTo
	ET    TextTranslationReqTo
	LT    TextTranslationReqTo
	HR    TextTranslationReqTo
	BS    TextTranslationReqTo
	KA    TextTranslationReqTo
	SL    TextTranslationReqTo
	CA    TextTranslationReqTo
	AF    TextTranslationReqTo
	MG    TextTranslationReqTo
	ID    TextTranslationReqTo
	FIL   TextTranslationReqTo
	SW    TextTranslationReqTo
	HU    TextTranslationReqTo
	SR    TextTranslationReqTo
	MK    TextTranslationReqTo
	UK    TextTranslationReqTo
	BG    TextTranslationReqTo
	MT    TextTranslationReqTo
	EL    TextTranslationReqTo
	IS    TextTranslationReqTo
	GA    TextTranslationReqTo
	CY    TextTranslationReqTo
	HT    TextTranslationReqTo
	NO    TextTranslationReqTo
	SV    TextTranslationReqTo
	FI    TextTranslationReqTo
	DA    TextTranslationReqTo
	ZH_TW TextTranslationReqTo
}

func GetTextTranslationReqToEnum() TextTranslationReqToEnum {
	return TextTranslationReqToEnum{
		ZH: TextTranslationReqTo{
			value: "zh",
		},
		EN: TextTranslationReqTo{
			value: "en",
		},
		JA: TextTranslationReqTo{
			value: "ja",
		},
		RU: TextTranslationReqTo{
			value: "ru",
		},
		ES: TextTranslationReqTo{
			value: "es",
		},
		DE: TextTranslationReqTo{
			value: "de",
		},
		AR: TextTranslationReqTo{
			value: "ar",
		},
		PT: TextTranslationReqTo{
			value: "pt",
		},
		TH: TextTranslationReqTo{
			value: "th",
		},
		TR: TextTranslationReqTo{
			value: "tr",
		},
		KO: TextTranslationReqTo{
			value: "ko",
		},
		FR: TextTranslationReqTo{
			value: "fr",
		},
		VI: TextTranslationReqTo{
			value: "vi",
		},
		HI: TextTranslationReqTo{
			value: "hi",
		},
		KM: TextTranslationReqTo{
			value: "km",
		},
		MY: TextTranslationReqTo{
			value: "my",
		},
		TA: TextTranslationReqTo{
			value: "ta",
		},
		FA: TextTranslationReqTo{
			value: "fa",
		},
		UR: TextTranslationReqTo{
			value: "ur",
		},
		BN: TextTranslationReqTo{
			value: "bn",
		},
		MR: TextTranslationReqTo{
			value: "mr",
		},
		GU: TextTranslationReqTo{
			value: "gu",
		},
		PA: TextTranslationReqTo{
			value: "pa",
		},
		TE: TextTranslationReqTo{
			value: "te",
		},
		HE: TextTranslationReqTo{
			value: "he",
		},
		CS: TextTranslationReqTo{
			value: "cs",
		},
		SK: TextTranslationReqTo{
			value: "sk",
		},
		RO: TextTranslationReqTo{
			value: "ro",
		},
		SQ: TextTranslationReqTo{
			value: "sq",
		},
		LV: TextTranslationReqTo{
			value: "lv",
		},
		ET: TextTranslationReqTo{
			value: "et",
		},
		LT: TextTranslationReqTo{
			value: "lt",
		},
		HR: TextTranslationReqTo{
			value: "hr",
		},
		BS: TextTranslationReqTo{
			value: "bs",
		},
		KA: TextTranslationReqTo{
			value: "ka",
		},
		SL: TextTranslationReqTo{
			value: "sl",
		},
		CA: TextTranslationReqTo{
			value: "ca",
		},
		AF: TextTranslationReqTo{
			value: "af",
		},
		MG: TextTranslationReqTo{
			value: "mg",
		},
		ID: TextTranslationReqTo{
			value: "id",
		},
		FIL: TextTranslationReqTo{
			value: "fil",
		},
		SW: TextTranslationReqTo{
			value: "sw",
		},
		HU: TextTranslationReqTo{
			value: "hu",
		},
		SR: TextTranslationReqTo{
			value: "sr",
		},
		MK: TextTranslationReqTo{
			value: "mk",
		},
		UK: TextTranslationReqTo{
			value: "uk",
		},
		BG: TextTranslationReqTo{
			value: "bg",
		},
		MT: TextTranslationReqTo{
			value: "mt",
		},
		EL: TextTranslationReqTo{
			value: "el",
		},
		IS: TextTranslationReqTo{
			value: "is",
		},
		GA: TextTranslationReqTo{
			value: "ga",
		},
		CY: TextTranslationReqTo{
			value: "cy",
		},
		HT: TextTranslationReqTo{
			value: "ht",
		},
		NO: TextTranslationReqTo{
			value: "no",
		},
		SV: TextTranslationReqTo{
			value: "sv",
		},
		FI: TextTranslationReqTo{
			value: "fi",
		},
		DA: TextTranslationReqTo{
			value: "da",
		},
		ZH_TW: TextTranslationReqTo{
			value: "zh-tw",
		},
	}
}

func (c TextTranslationReqTo) Value() string {
	return c.value
}

func (c TextTranslationReqTo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextTranslationReqTo) UnmarshalJSON(b []byte) error {
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

type TextTranslationReqScene struct {
	value string
}

type TextTranslationReqSceneEnum struct {
	COMMON TextTranslationReqScene
}

func GetTextTranslationReqSceneEnum() TextTranslationReqSceneEnum {
	return TextTranslationReqSceneEnum{
		COMMON: TextTranslationReqScene{
			value: "common",
		},
	}
}

func (c TextTranslationReqScene) Value() string {
	return c.value
}

func (c TextTranslationReqScene) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextTranslationReqScene) UnmarshalJSON(b []byte) error {
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
