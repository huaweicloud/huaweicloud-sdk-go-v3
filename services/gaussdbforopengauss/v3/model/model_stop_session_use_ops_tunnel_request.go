package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopSessionUseOpsTunnelRequest Request Object
type StopSessionUseOpsTunnelRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *StopSessionUseOpsTunnelRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *KillSessionRequestBody `json:"body,omitempty"`
}

func (o StopSessionUseOpsTunnelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSessionUseOpsTunnelRequest struct{}"
	}

	return strings.Join([]string{"StopSessionUseOpsTunnelRequest", string(data)}, " ")
}

type StopSessionUseOpsTunnelRequestXLanguage struct {
	value string
}

type StopSessionUseOpsTunnelRequestXLanguageEnum struct {
	ZH_CN StopSessionUseOpsTunnelRequestXLanguage
	EN_US StopSessionUseOpsTunnelRequestXLanguage
}

func GetStopSessionUseOpsTunnelRequestXLanguageEnum() StopSessionUseOpsTunnelRequestXLanguageEnum {
	return StopSessionUseOpsTunnelRequestXLanguageEnum{
		ZH_CN: StopSessionUseOpsTunnelRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopSessionUseOpsTunnelRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopSessionUseOpsTunnelRequestXLanguage) Value() string {
	return c.value
}

func (c StopSessionUseOpsTunnelRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopSessionUseOpsTunnelRequestXLanguage) UnmarshalJSON(b []byte) error {
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
