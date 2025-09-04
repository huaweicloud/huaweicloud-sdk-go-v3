package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeFullDeadLockSwitchRequest Request Object
type ChangeFullDeadLockSwitchRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ChangeFullDeadLockSwitchRequestXLanguage `json:"X-Language,omitempty"`

	Body *ChangeFullDeadLockSwitchRequestBody `json:"body,omitempty"`
}

func (o ChangeFullDeadLockSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFullDeadLockSwitchRequest struct{}"
	}

	return strings.Join([]string{"ChangeFullDeadLockSwitchRequest", string(data)}, " ")
}

type ChangeFullDeadLockSwitchRequestXLanguage struct {
	value string
}

type ChangeFullDeadLockSwitchRequestXLanguageEnum struct {
	ZH_CN ChangeFullDeadLockSwitchRequestXLanguage
	EN_US ChangeFullDeadLockSwitchRequestXLanguage
}

func GetChangeFullDeadLockSwitchRequestXLanguageEnum() ChangeFullDeadLockSwitchRequestXLanguageEnum {
	return ChangeFullDeadLockSwitchRequestXLanguageEnum{
		ZH_CN: ChangeFullDeadLockSwitchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeFullDeadLockSwitchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeFullDeadLockSwitchRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeFullDeadLockSwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeFullDeadLockSwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
