package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 文本相似度请求体
type TextSimilarityRequest struct {

	// 待计算文本1，中文长度1~512，英文长度1~2000，文本编码为UTF-8。
	Text1 string `json:"text1" xml:"text1"`

	// 待计算文本2，中文长度1~512，英文长度1~2000，文本编码为UTF-8。
	Text2 string `json:"text2" xml:"text2"`

	// 支持的文本语言类型，目前支持中文（zh）和英文（en），默认为中文。
	Lang *TextSimilarityRequestLang `json:"lang,omitempty" xml:"lang"`
}

func (o TextSimilarityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextSimilarityRequest struct{}"
	}

	return strings.Join([]string{"TextSimilarityRequest", string(data)}, " ")
}

type TextSimilarityRequestLang struct {
	value string
}

type TextSimilarityRequestLangEnum struct {
	ZH TextSimilarityRequestLang
	EN TextSimilarityRequestLang
}

func GetTextSimilarityRequestLangEnum() TextSimilarityRequestLangEnum {
	return TextSimilarityRequestLangEnum{
		ZH: TextSimilarityRequestLang{
			value: "zh",
		},
		EN: TextSimilarityRequestLang{
			value: "en",
		},
	}
}

func (c TextSimilarityRequestLang) Value() string {
	return c.value
}

func (c TextSimilarityRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextSimilarityRequestLang) UnmarshalJSON(b []byte) error {
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
