package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type KeywordExtractReq struct {

	// 待分析文本，长度为1~512，文本编码为UTF-8。
	Text string `json:"text" xml:"text"`

	// 返回关键词的最大数量，默认为5。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 支持的文本语言类型，目前只支持中文，默认为zh。
	Lang *KeywordExtractReqLang `json:"lang,omitempty" xml:"lang"`
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

func (c KeywordExtractReqLang) Value() string {
	return c.value
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
