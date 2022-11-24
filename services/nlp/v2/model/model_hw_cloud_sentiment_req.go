package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type HwCloudSentimentReq struct {

	// 待分析文本。文本编码要求为utf-8。仅支持中文和英文情感分析。 限定180个字符以内，超过180个字符，只截取前180个字符。
	Content string `json:"content"`

	// 支持的文本语言类型，目前支持中文（zh）和英文（en），默认为中文。
	Lang *HwCloudSentimentReqLang `json:"lang,omitempty"`
}

func (o HwCloudSentimentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwCloudSentimentReq struct{}"
	}

	return strings.Join([]string{"HwCloudSentimentReq", string(data)}, " ")
}

type HwCloudSentimentReqLang struct {
	value string
}

type HwCloudSentimentReqLangEnum struct {
	EN HwCloudSentimentReqLang
	ZH HwCloudSentimentReqLang
}

func GetHwCloudSentimentReqLangEnum() HwCloudSentimentReqLangEnum {
	return HwCloudSentimentReqLangEnum{
		EN: HwCloudSentimentReqLang{
			value: "en",
		},
		ZH: HwCloudSentimentReqLang{
			value: "zh",
		},
	}
}

func (c HwCloudSentimentReqLang) Value() string {
	return c.value
}

func (c HwCloudSentimentReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HwCloudSentimentReqLang) UnmarshalJSON(b []byte) error {
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
