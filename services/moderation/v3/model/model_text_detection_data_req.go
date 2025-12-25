package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TextDetectionDataReq struct {

	// 待检测文本，编码格式为“utf-8”，限定10000个字符以内，文本长度超过10000个字符时，只检测前10000个字符。
	Text string `json:"text"`

	// 支持检测的文本语言
	Language *TextDetectionDataReqLanguage `json:"language,omitempty"`
}

func (o TextDetectionDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionDataReq struct{}"
	}

	return strings.Join([]string{"TextDetectionDataReq", string(data)}, " ")
}

type TextDetectionDataReqLanguage struct {
	value string
}

type TextDetectionDataReqLanguageEnum struct {
	ZH TextDetectionDataReqLanguage
}

func GetTextDetectionDataReqLanguageEnum() TextDetectionDataReqLanguageEnum {
	return TextDetectionDataReqLanguageEnum{
		ZH: TextDetectionDataReqLanguage{
			value: "zh",
		},
	}
}

func (c TextDetectionDataReqLanguage) Value() string {
	return c.value
}

func (c TextDetectionDataReqLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TextDetectionDataReqLanguage) UnmarshalJSON(b []byte) error {
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
