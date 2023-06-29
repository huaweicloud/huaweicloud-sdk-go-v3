package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AspectSentimentRequest
type AspectSentimentRequest struct {

	// 待分析文本。仅支持中文，文本编码要求为utf-8。 建议文本长度1000个字符以内。
	Content string `json:"content"`

	// 取值如下： 1 手机领域
	Type AspectSentimentRequestType `json:"type"`
}

func (o AspectSentimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AspectSentimentRequest struct{}"
	}

	return strings.Join([]string{"AspectSentimentRequest", string(data)}, " ")
}

type AspectSentimentRequestType struct {
	value int32
}

type AspectSentimentRequestTypeEnum struct {
	E_1 AspectSentimentRequestType
}

func GetAspectSentimentRequestTypeEnum() AspectSentimentRequestTypeEnum {
	return AspectSentimentRequestTypeEnum{
		E_1: AspectSentimentRequestType{
			value: 1,
		},
	}
}

func (c AspectSentimentRequestType) Value() int32 {
	return c.value
}

func (c AspectSentimentRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AspectSentimentRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
