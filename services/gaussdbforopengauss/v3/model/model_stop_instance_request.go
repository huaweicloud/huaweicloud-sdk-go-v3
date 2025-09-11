package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopInstanceRequest Request Object
type StopInstanceRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *StopInstanceRequestXLanguage `json:"X-Language,omitempty"`

	// 需要停止的实例的ID
	InstanceId string `json:"instance_id"`

	Body *StopInstanceRequestBody `json:"body,omitempty"`
}

func (o StopInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceRequest", string(data)}, " ")
}

type StopInstanceRequestXLanguage struct {
	value string
}

type StopInstanceRequestXLanguageEnum struct {
	ZH_CN StopInstanceRequestXLanguage
	EN_US StopInstanceRequestXLanguage
}

func GetStopInstanceRequestXLanguageEnum() StopInstanceRequestXLanguageEnum {
	return StopInstanceRequestXLanguageEnum{
		ZH_CN: StopInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopInstanceRequestXLanguage) Value() string {
	return c.value
}

func (c StopInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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
