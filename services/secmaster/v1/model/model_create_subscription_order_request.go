package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSubscriptionOrderRequest Request Object
type CreateSubscriptionOrderRequest struct {

	// 用户当前语言环境 zh-cn or en-us.
	XLanguage CreateSubscriptionOrderRequestXLanguage `json:"X-Language"`

	Body *OrderInfoReq `json:"body,omitempty"`
}

func (o CreateSubscriptionOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionOrderRequest", string(data)}, " ")
}

type CreateSubscriptionOrderRequestXLanguage struct {
	value string
}

type CreateSubscriptionOrderRequestXLanguageEnum struct {
	ZH_CN CreateSubscriptionOrderRequestXLanguage
	EN_US CreateSubscriptionOrderRequestXLanguage
}

func GetCreateSubscriptionOrderRequestXLanguageEnum() CreateSubscriptionOrderRequestXLanguageEnum {
	return CreateSubscriptionOrderRequestXLanguageEnum{
		ZH_CN: CreateSubscriptionOrderRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateSubscriptionOrderRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateSubscriptionOrderRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSubscriptionOrderRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubscriptionOrderRequestXLanguage) UnmarshalJSON(b []byte) error {
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
