package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateCompareTaskRequest struct {
	// 请求语言类型

	XLanguage CreateCompareTaskRequestXLanguage `json:"X-Language"`

	Body *CreateCompareTaskReq `json:"body,omitempty"`
}

func (o CreateCompareTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskRequest", string(data)}, " ")
}

type CreateCompareTaskRequestXLanguage struct {
	value string
}

type CreateCompareTaskRequestXLanguageEnum struct {
	EN_US CreateCompareTaskRequestXLanguage
	ZH_CN CreateCompareTaskRequestXLanguage
}

func GetCreateCompareTaskRequestXLanguageEnum() CreateCompareTaskRequestXLanguageEnum {
	return CreateCompareTaskRequestXLanguageEnum{
		EN_US: CreateCompareTaskRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateCompareTaskRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateCompareTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCompareTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
