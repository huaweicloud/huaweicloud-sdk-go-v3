package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlLimitTaskRequest Request Object
type CreateSqlLimitTaskRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *CreateSqlLimitTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateSqlLimitTaskRequestBody `json:"body,omitempty"`
}

func (o CreateSqlLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitTaskRequest", string(data)}, " ")
}

type CreateSqlLimitTaskRequestXLanguage struct {
	value string
}

type CreateSqlLimitTaskRequestXLanguageEnum struct {
	ZH_CN CreateSqlLimitTaskRequestXLanguage
	EN_US CreateSqlLimitTaskRequestXLanguage
}

func GetCreateSqlLimitTaskRequestXLanguageEnum() CreateSqlLimitTaskRequestXLanguageEnum {
	return CreateSqlLimitTaskRequestXLanguageEnum{
		ZH_CN: CreateSqlLimitTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateSqlLimitTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateSqlLimitTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSqlLimitTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlLimitTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
