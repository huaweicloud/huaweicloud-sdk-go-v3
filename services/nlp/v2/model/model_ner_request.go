package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 命名实体识别post请求体
type NerRequest struct {

	// 待分析文本，中文长度为1~512，英文和西班牙文长度为1~2000，文本编码为UTF-8。
	Text string `json:"text"`

	// 支持的文本语言类型，目前支持中文（zh）,英文（en）,和西班牙文（es），默认为中文。
	Lang *NerRequestLang `json:"lang,omitempty"`
}

func (o NerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NerRequest struct{}"
	}

	return strings.Join([]string{"NerRequest", string(data)}, " ")
}

type NerRequestLang struct {
	value string
}

type NerRequestLangEnum struct {
	ZH NerRequestLang
	EN NerRequestLang
	ES NerRequestLang
}

func GetNerRequestLangEnum() NerRequestLangEnum {
	return NerRequestLangEnum{
		ZH: NerRequestLang{
			value: "zh",
		},
		EN: NerRequestLang{
			value: "en",
		},
		ES: NerRequestLang{
			value: "es",
		},
	}
}

func (c NerRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NerRequestLang) UnmarshalJSON(b []byte) error {
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
