package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NotificationMode 发送通知的方式。  - 1: 短信  - 2: 企业微信  - 3: 钉钉  - 4: 飞书  - 5: 电话
type NotificationMode struct {
	value string
}

type NotificationModeEnum struct {
	E_1 NotificationMode
	E_2 NotificationMode
	E_3 NotificationMode
	E_4 NotificationMode
	E_5 NotificationMode
}

func GetNotificationModeEnum() NotificationModeEnum {
	return NotificationModeEnum{
		E_1: NotificationMode{
			value: "1",
		},
		E_2: NotificationMode{
			value: "2",
		},
		E_3: NotificationMode{
			value: "3",
		},
		E_4: NotificationMode{
			value: "4",
		},
		E_5: NotificationMode{
			value: "5",
		},
	}
}

func (c NotificationMode) Value() string {
	return c.value
}

func (c NotificationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationMode) UnmarshalJSON(b []byte) error {
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
