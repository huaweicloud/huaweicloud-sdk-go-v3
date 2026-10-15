package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeBigKeyHotKeySwitchRequest Request Object
type ChangeBigKeyHotKeySwitchRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。zh-cn：中文。
	XLanguage *ChangeBigKeyHotKeySwitchRequestXLanguage `json:"X-Language,omitempty"`

	Body *BigKeyHotKeySwitchReq `json:"body,omitempty"`
}

func (o ChangeBigKeyHotKeySwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBigKeyHotKeySwitchRequest struct{}"
	}

	return strings.Join([]string{"ChangeBigKeyHotKeySwitchRequest", string(data)}, " ")
}

type ChangeBigKeyHotKeySwitchRequestXLanguage struct {
	value string
}

type ChangeBigKeyHotKeySwitchRequestXLanguageEnum struct {
	ZH_CN ChangeBigKeyHotKeySwitchRequestXLanguage
	EN_US ChangeBigKeyHotKeySwitchRequestXLanguage
}

func GetChangeBigKeyHotKeySwitchRequestXLanguageEnum() ChangeBigKeyHotKeySwitchRequestXLanguageEnum {
	return ChangeBigKeyHotKeySwitchRequestXLanguageEnum{
		ZH_CN: ChangeBigKeyHotKeySwitchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeBigKeyHotKeySwitchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeBigKeyHotKeySwitchRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeBigKeyHotKeySwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeBigKeyHotKeySwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
