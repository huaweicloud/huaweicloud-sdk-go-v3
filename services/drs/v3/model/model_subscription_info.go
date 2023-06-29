package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SubscriptionInfo 收件方式与信息体
type SubscriptionInfo struct {

	// 短信或者邮件列表
	Endpoints *[]string `json:"endpoints,omitempty"`

	// 收件方式，sms：短信,email：邮件
	Protocol *SubscriptionInfoProtocol `json:"protocol,omitempty"`
}

func (o SubscriptionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionInfo", string(data)}, " ")
}

type SubscriptionInfoProtocol struct {
	value string
}

type SubscriptionInfoProtocolEnum struct {
	SMS   SubscriptionInfoProtocol
	EMAIL SubscriptionInfoProtocol
}

func GetSubscriptionInfoProtocolEnum() SubscriptionInfoProtocolEnum {
	return SubscriptionInfoProtocolEnum{
		SMS: SubscriptionInfoProtocol{
			value: "sms",
		},
		EMAIL: SubscriptionInfoProtocol{
			value: "email",
		},
	}
}

func (c SubscriptionInfoProtocol) Value() string {
	return c.value
}

func (c SubscriptionInfoProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionInfoProtocol) UnmarshalJSON(b []byte) error {
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
