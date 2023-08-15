package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AspectSentimentAdvanceRequest
type AspectSentimentAdvanceRequest struct {

	// 待分析文本。文本编码要求为utf-8，仅支持中文。 限定4096个字符以内，建议文本长度300个字符以内。
	Content string `json:"content"`

	// 取值如下： 1 手机领域 2 汽车领域
	Type AspectSentimentAdvanceRequestType `json:"type"`
}

func (o AspectSentimentAdvanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AspectSentimentAdvanceRequest struct{}"
	}

	return strings.Join([]string{"AspectSentimentAdvanceRequest", string(data)}, " ")
}

type AspectSentimentAdvanceRequestType struct {
	value int32
}

type AspectSentimentAdvanceRequestTypeEnum struct {
	E_1 AspectSentimentAdvanceRequestType
	E_2 AspectSentimentAdvanceRequestType
}

func GetAspectSentimentAdvanceRequestTypeEnum() AspectSentimentAdvanceRequestTypeEnum {
	return AspectSentimentAdvanceRequestTypeEnum{
		E_1: AspectSentimentAdvanceRequestType{
			value: 1,
		}, E_2: AspectSentimentAdvanceRequestType{
			value: 2,
		},
	}
}

func (c AspectSentimentAdvanceRequestType) Value() int32 {
	return c.value
}

func (c AspectSentimentAdvanceRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AspectSentimentAdvanceRequestType) UnmarshalJSON(b []byte) error {
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
