package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSqlLimitTaskRequest Request Object
type UpdateSqlLimitTaskRequest struct {

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 限流任务ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *UpdateSqlLimitTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateSqlLimitTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateSqlLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitTaskRequest", string(data)}, " ")
}

type UpdateSqlLimitTaskRequestXLanguage struct {
	value string
}

type UpdateSqlLimitTaskRequestXLanguageEnum struct {
	ZH_CN UpdateSqlLimitTaskRequestXLanguage
	EN_US UpdateSqlLimitTaskRequestXLanguage
}

func GetUpdateSqlLimitTaskRequestXLanguageEnum() UpdateSqlLimitTaskRequestXLanguageEnum {
	return UpdateSqlLimitTaskRequestXLanguageEnum{
		ZH_CN: UpdateSqlLimitTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateSqlLimitTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateSqlLimitTaskRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateSqlLimitTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlLimitTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
