package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type IntentReq struct {
	// 支持的文本语言类型，目前只支持中文，默认为zh。

	Lang *IntentReqLang `json:"lang,omitempty"`
	// 待分析文本列表，UTF-8编码，限定32个字符以内，文本长度超过32个字符时，只检测前32个字符。

	Text string `json:"text"`
}

func (o IntentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntentReq struct{}"
	}

	return strings.Join([]string{"IntentReq", string(data)}, " ")
}

type IntentReqLang struct {
	value string
}

type IntentReqLangEnum struct {
	ZH IntentReqLang
}

func GetIntentReqLangEnum() IntentReqLangEnum {
	return IntentReqLangEnum{
		ZH: IntentReqLang{
			value: "zh",
		},
	}
}

func (c IntentReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IntentReqLang) UnmarshalJSON(b []byte) error {
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
