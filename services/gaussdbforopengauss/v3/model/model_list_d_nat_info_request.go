package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDNatInfoRequest Request Object
type ListDNatInfoRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListDNatInfoRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`
}

func (o ListDNatInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDNatInfoRequest struct{}"
	}

	return strings.Join([]string{"ListDNatInfoRequest", string(data)}, " ")
}

type ListDNatInfoRequestXLanguage struct {
	value string
}

type ListDNatInfoRequestXLanguageEnum struct {
	ZH_CN ListDNatInfoRequestXLanguage
	EN_US ListDNatInfoRequestXLanguage
}

func GetListDNatInfoRequestXLanguageEnum() ListDNatInfoRequestXLanguageEnum {
	return ListDNatInfoRequestXLanguageEnum{
		ZH_CN: ListDNatInfoRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDNatInfoRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDNatInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ListDNatInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDNatInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
