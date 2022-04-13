package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 请求消息
type DomainSentimentReq struct {
	// 待分析文本。文本编码要求为utf-8。仅支持中文情感分析。 type为1（电商领域评论）时，限定200个字符以内，文本长度超过200个字符时，只检测前200个字符。 type为2（汽车领域评论）时，限定400个字符以内，文本长度超过400个字符时，只检测前400个字符。

	Content string `json:"content"`
	// 取值如下： 0：自适应领域，根据输入内容自动识别适应领域。 1：电商领域，适用于电商领域评论。 2：汽车领域，适用于汽车领域评论。

	Type *DomainSentimentReqType `json:"type,omitempty"`
}

func (o DomainSentimentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSentimentReq struct{}"
	}

	return strings.Join([]string{"DomainSentimentReq", string(data)}, " ")
}

type DomainSentimentReqType struct {
	value int32
}

type DomainSentimentReqTypeEnum struct {
	E_0 DomainSentimentReqType
	E_1 DomainSentimentReqType
	E_2 DomainSentimentReqType
}

func GetDomainSentimentReqTypeEnum() DomainSentimentReqTypeEnum {
	return DomainSentimentReqTypeEnum{
		E_0: DomainSentimentReqType{
			value: 0,
		}, E_1: DomainSentimentReqType{
			value: 1,
		}, E_2: DomainSentimentReqType{
			value: 2,
		},
	}
}

func (c DomainSentimentReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainSentimentReqType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
