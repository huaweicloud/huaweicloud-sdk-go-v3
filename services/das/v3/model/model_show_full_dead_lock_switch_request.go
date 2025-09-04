package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFullDeadLockSwitchRequest Request Object
type ShowFullDeadLockSwitchRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ShowFullDeadLockSwitchRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowFullDeadLockSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullDeadLockSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowFullDeadLockSwitchRequest", string(data)}, " ")
}

type ShowFullDeadLockSwitchRequestXLanguage struct {
	value string
}

type ShowFullDeadLockSwitchRequestXLanguageEnum struct {
	ZH_CN ShowFullDeadLockSwitchRequestXLanguage
	EN_US ShowFullDeadLockSwitchRequestXLanguage
}

func GetShowFullDeadLockSwitchRequestXLanguageEnum() ShowFullDeadLockSwitchRequestXLanguageEnum {
	return ShowFullDeadLockSwitchRequestXLanguageEnum{
		ZH_CN: ShowFullDeadLockSwitchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowFullDeadLockSwitchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowFullDeadLockSwitchRequestXLanguage) Value() string {
	return c.value
}

func (c ShowFullDeadLockSwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFullDeadLockSwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
