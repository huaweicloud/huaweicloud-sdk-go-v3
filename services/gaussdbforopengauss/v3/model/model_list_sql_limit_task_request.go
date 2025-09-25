package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlLimitTaskRequest Request Object
type ListSqlLimitTaskRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListSqlLimitTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *ListSqlLimitTaskRequestBody `json:"body,omitempty"`
}

func (o ListSqlLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"ListSqlLimitTaskRequest", string(data)}, " ")
}

type ListSqlLimitTaskRequestXLanguage struct {
	value string
}

type ListSqlLimitTaskRequestXLanguageEnum struct {
	ZH_CN ListSqlLimitTaskRequestXLanguage
	EN_US ListSqlLimitTaskRequestXLanguage
}

func GetListSqlLimitTaskRequestXLanguageEnum() ListSqlLimitTaskRequestXLanguageEnum {
	return ListSqlLimitTaskRequestXLanguageEnum{
		ZH_CN: ListSqlLimitTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSqlLimitTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSqlLimitTaskRequestXLanguage) Value() string {
	return c.value
}

func (c ListSqlLimitTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlLimitTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
