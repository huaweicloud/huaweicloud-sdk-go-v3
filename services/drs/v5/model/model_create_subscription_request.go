package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSubscriptionRequest Request Object
type CreateSubscriptionRequest struct {

	// 请求语言类型。
	XLanguage *CreateSubscriptionRequestXLanguage `json:"X-Language,omitempty"`

	Body *SingleCreateSubscriptionReq `json:"body,omitempty"`
}

func (o CreateSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionRequest", string(data)}, " ")
}

type CreateSubscriptionRequestXLanguage struct {
	value string
}

type CreateSubscriptionRequestXLanguageEnum struct {
	EN_US CreateSubscriptionRequestXLanguage
	ZH_CN CreateSubscriptionRequestXLanguage
}

func GetCreateSubscriptionRequestXLanguageEnum() CreateSubscriptionRequestXLanguageEnum {
	return CreateSubscriptionRequestXLanguageEnum{
		EN_US: CreateSubscriptionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateSubscriptionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateSubscriptionRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSubscriptionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubscriptionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
