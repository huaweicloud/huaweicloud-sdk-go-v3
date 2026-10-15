package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBigKeyHotKeySwitchRequest Request Object
type ShowBigKeyHotKeySwitchRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。zh-cn：中文。
	XLanguage *ShowBigKeyHotKeySwitchRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowBigKeyHotKeySwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigKeyHotKeySwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowBigKeyHotKeySwitchRequest", string(data)}, " ")
}

type ShowBigKeyHotKeySwitchRequestXLanguage struct {
	value string
}

type ShowBigKeyHotKeySwitchRequestXLanguageEnum struct {
	ZH_CN ShowBigKeyHotKeySwitchRequestXLanguage
	EN_US ShowBigKeyHotKeySwitchRequestXLanguage
}

func GetShowBigKeyHotKeySwitchRequestXLanguageEnum() ShowBigKeyHotKeySwitchRequestXLanguageEnum {
	return ShowBigKeyHotKeySwitchRequestXLanguageEnum{
		ZH_CN: ShowBigKeyHotKeySwitchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowBigKeyHotKeySwitchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowBigKeyHotKeySwitchRequestXLanguage) Value() string {
	return c.value
}

func (c ShowBigKeyHotKeySwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBigKeyHotKeySwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
