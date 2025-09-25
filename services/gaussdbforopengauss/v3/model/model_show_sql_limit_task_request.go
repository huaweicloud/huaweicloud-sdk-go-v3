package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlLimitTaskRequest Request Object
type ShowSqlLimitTaskRequest struct {

	// **参数解释**: 限流任务ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowSqlLimitTaskRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowSqlLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitTaskRequest", string(data)}, " ")
}

type ShowSqlLimitTaskRequestXLanguage struct {
	value string
}

type ShowSqlLimitTaskRequestXLanguageEnum struct {
	ZH_CN ShowSqlLimitTaskRequestXLanguage
	EN_US ShowSqlLimitTaskRequestXLanguage
}

func GetShowSqlLimitTaskRequestXLanguageEnum() ShowSqlLimitTaskRequestXLanguageEnum {
	return ShowSqlLimitTaskRequestXLanguageEnum{
		ZH_CN: ShowSqlLimitTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSqlLimitTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSqlLimitTaskRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSqlLimitTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlLimitTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
