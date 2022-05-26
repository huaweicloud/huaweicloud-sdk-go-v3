package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 文本相似度请求体
type TextSimilarityAdvanceRequest struct {

	// 待计算文本1，长度1~512，文本编码为UTF-8。
	Text1 string `json:"text1"`

	// 待计算文本1，长度1~512，文本编码为UTF-8。
	Text2 string `json:"text2"`

	// 支持的文本语言类型，目前只支持中文（zh），默认为中文。
	Lang *TextSimilarityAdvanceRequestLang `json:"lang,omitempty"`
}

func (o TextSimilarityAdvanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextSimilarityAdvanceRequest struct{}"
	}

	return strings.Join([]string{"TextSimilarityAdvanceRequest", string(data)}, " ")
}

type TextSimilarityAdvanceRequestLang struct {
	value string
}

type TextSimilarityAdvanceRequestLangEnum struct {
	ZH TextSimilarityAdvanceRequestLang
}

func GetTextSimilarityAdvanceRequestLangEnum() TextSimilarityAdvanceRequestLangEnum {
	return TextSimilarityAdvanceRequestLangEnum{
		ZH: TextSimilarityAdvanceRequestLang{
			value: "zh",
		},
	}
}

func (c TextSimilarityAdvanceRequestLang) Value() string {
	return c.value
}

func (c TextSimilarityAdvanceRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextSimilarityAdvanceRequestLang) UnmarshalJSON(b []byte) error {
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
