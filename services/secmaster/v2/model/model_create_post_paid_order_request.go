package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePostPaidOrderRequest Request Object
type CreatePostPaidOrderRequest struct {

	// 用户当前语言环境
	XLanguage CreatePostPaidOrderRequestXLanguage `json:"X-Language"`

	Body *PostPaidParam `json:"body,omitempty"`
}

func (o CreatePostPaidOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidOrderRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidOrderRequest", string(data)}, " ")
}

type CreatePostPaidOrderRequestXLanguage struct {
	value string
}

type CreatePostPaidOrderRequestXLanguageEnum struct {
	ZH_CN CreatePostPaidOrderRequestXLanguage
	EN_US CreatePostPaidOrderRequestXLanguage
}

func GetCreatePostPaidOrderRequestXLanguageEnum() CreatePostPaidOrderRequestXLanguageEnum {
	return CreatePostPaidOrderRequestXLanguageEnum{
		ZH_CN: CreatePostPaidOrderRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreatePostPaidOrderRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreatePostPaidOrderRequestXLanguage) Value() string {
	return c.value
}

func (c CreatePostPaidOrderRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidOrderRequestXLanguage) UnmarshalJSON(b []byte) error {
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
