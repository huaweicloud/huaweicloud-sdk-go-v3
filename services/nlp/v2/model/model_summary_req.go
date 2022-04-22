package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SummaryReq struct {

	// 文本正文（目前仅支持UTF-8编码），长度不超过10000字。
	Content string `json:"content"`

	// 支持的文本语言类型，目前支持中文（zh）和英文（en），默认为中文。
	Lang *SummaryReqLang `json:"lang,omitempty"`

	// 生成摘要的长度限制。length_limit > 1，则返回结果为字数不小于该值且最接近该值的摘要。 0 <= length_limit <= 1，则返回结果为长度百分比不小于该值且最接近该值的摘要。 默认数值为0.3。
	LengthLimit *float32 `json:"length_limit,omitempty"`

	// 文本标题（目前仅支持UTF-8编码），长度不超过1000字。
	Title *string `json:"title,omitempty"`
}

func (o SummaryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummaryReq struct{}"
	}

	return strings.Join([]string{"SummaryReq", string(data)}, " ")
}

type SummaryReqLang struct {
	value string
}

type SummaryReqLangEnum struct {
	ZH SummaryReqLang
	EN SummaryReqLang
}

func GetSummaryReqLangEnum() SummaryReqLangEnum {
	return SummaryReqLangEnum{
		ZH: SummaryReqLang{
			value: "zh",
		},
		EN: SummaryReqLang{
			value: "en",
		},
	}
}

func (c SummaryReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SummaryReqLang) UnmarshalJSON(b []byte) error {
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
