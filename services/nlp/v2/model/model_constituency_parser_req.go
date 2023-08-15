package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConstituencyParserReq struct {

	// 支持的文本语言类型，目前支持中文（zh）。
	Lang *ConstituencyParserReqLang `json:"lang,omitempty"`

	// 待分析文本，长度为1~32。
	Text string `json:"text"`
}

func (o ConstituencyParserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConstituencyParserReq struct{}"
	}

	return strings.Join([]string{"ConstituencyParserReq", string(data)}, " ")
}

type ConstituencyParserReqLang struct {
	value string
}

type ConstituencyParserReqLangEnum struct {
	ZH ConstituencyParserReqLang
}

func GetConstituencyParserReqLangEnum() ConstituencyParserReqLangEnum {
	return ConstituencyParserReqLangEnum{
		ZH: ConstituencyParserReqLang{
			value: "zh",
		},
	}
}

func (c ConstituencyParserReqLang) Value() string {
	return c.value
}

func (c ConstituencyParserReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConstituencyParserReqLang) UnmarshalJSON(b []byte) error {
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
