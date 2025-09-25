package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSqlLimitTaskRequest Request Object
type DeleteSqlLimitTaskRequest struct {

	// **参数解释**: 限流任务ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *DeleteSqlLimitTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteSqlLimitTaskRequestBody `json:"body,omitempty"`
}

func (o DeleteSqlLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitTaskRequest", string(data)}, " ")
}

type DeleteSqlLimitTaskRequestXLanguage struct {
	value string
}

type DeleteSqlLimitTaskRequestXLanguageEnum struct {
	ZH_CN DeleteSqlLimitTaskRequestXLanguage
	EN_US DeleteSqlLimitTaskRequestXLanguage
}

func GetDeleteSqlLimitTaskRequestXLanguageEnum() DeleteSqlLimitTaskRequestXLanguageEnum {
	return DeleteSqlLimitTaskRequestXLanguageEnum{
		ZH_CN: DeleteSqlLimitTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteSqlLimitTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteSqlLimitTaskRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSqlLimitTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSqlLimitTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
