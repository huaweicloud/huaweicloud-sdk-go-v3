package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// This is a auto create Body Object
type KeywordExtractReq struct {
	// 待分析文本，长度为1~512，文本编码为UTF-8。

	Text string `json:"text"`
	// 返回关键词的最大数量，默认为5。

	Limit *int32 `json:"limit,omitempty"`
	// 支持的文本语言类型，目前只支持中文，默认为zh。

	Lang *KeywordExtractReqLang `json:"lang,omitempty"`
}

func (o KeywordExtractReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordExtractReq struct{}"
	}

	return strings.Join([]string{"KeywordExtractReq", string(data)}, " ")
}

type KeywordExtractReqLang struct {
	value string
}

type KeywordExtractReqLangEnum struct {
	ZH KeywordExtractReqLang
}

func GetKeywordExtractReqLangEnum() KeywordExtractReqLangEnum {
	return KeywordExtractReqLangEnum{
		ZH: KeywordExtractReqLang{
			value: "zh",
		},
	}
}

func (c KeywordExtractReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KeywordExtractReqLang) UnmarshalJSON(b []byte) error {
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
