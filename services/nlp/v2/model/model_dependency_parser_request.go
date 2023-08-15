package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DependencyParserRequest 依存句法分析请求体
type DependencyParserRequest struct {

	// 待分析文本，长度为1~32，文本编码为utf-8。
	Text string `json:"text"`

	// 支持的文本语言类型，目前只支持中文，默认为zh。
	Lang *DependencyParserRequestLang `json:"lang,omitempty"`
}

func (o DependencyParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependencyParserRequest struct{}"
	}

	return strings.Join([]string{"DependencyParserRequest", string(data)}, " ")
}

type DependencyParserRequestLang struct {
	value string
}

type DependencyParserRequestLangEnum struct {
	ZH DependencyParserRequestLang
}

func GetDependencyParserRequestLangEnum() DependencyParserRequestLangEnum {
	return DependencyParserRequestLangEnum{
		ZH: DependencyParserRequestLang{
			value: "zh",
		},
	}
}

func (c DependencyParserRequestLang) Value() string {
	return c.value
}

func (c DependencyParserRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DependencyParserRequestLang) UnmarshalJSON(b []byte) error {
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
