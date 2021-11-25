package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type TextTranslationReq struct {
	// 待翻译文本，仅支持utf-8编码，长度不超过5000字符。

	Text string `json:"text"`
	// 翻译原语言，具体取值见支持的语言列表: zh    中文 en    英文 ru    俄语 ja    日文 de    德文 fr    法文 es    西班牙文 ko    韩语 auto  自动检测输入语种并翻译成目标语种，您需要指定目标语种。

	From TextTranslationReqFrom `json:"from"`
	// 翻译原语言，具体取值见支持的语言列表: zh    中文 en    英文 ru    俄语 ja    日文 de    德文 fr    法文 es    西班牙文 ko    韩语

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
	ZH   TextTranslationReqFrom
	EN   TextTranslationReqFrom
	JA   TextTranslationReqFrom
	RU   TextTranslationReqFrom
	KO   TextTranslationReqFrom
	FR   TextTranslationReqFrom
	ES   TextTranslationReqFrom
	DE   TextTranslationReqFrom
	AUTO TextTranslationReqFrom
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
		KO: TextTranslationReqFrom{
			value: "ko",
		},
		FR: TextTranslationReqFrom{
			value: "fr",
		},
		ES: TextTranslationReqFrom{
			value: "es",
		},
		DE: TextTranslationReqFrom{
			value: "de",
		},
		AUTO: TextTranslationReqFrom{
			value: "auto",
		},
	}
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
	ZH TextTranslationReqTo
	EN TextTranslationReqTo
	JA TextTranslationReqTo
	RU TextTranslationReqTo
	KO TextTranslationReqTo
	FR TextTranslationReqTo
	ES TextTranslationReqTo
	DE TextTranslationReqTo
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
		KO: TextTranslationReqTo{
			value: "ko",
		},
		FR: TextTranslationReqTo{
			value: "fr",
		},
		ES: TextTranslationReqTo{
			value: "es",
		},
		DE: TextTranslationReqTo{
			value: "de",
		},
	}
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
