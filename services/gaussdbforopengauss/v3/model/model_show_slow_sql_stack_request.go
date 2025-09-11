package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSlowSqlStackRequest Request Object
type ShowSlowSqlStackRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *ShowSlowSqlStackRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 线程ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Pid string `json:"pid"`

	// **参数解释**: 节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 节点上的组件ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	CompId *string `json:"comp_id,omitempty"`
}

func (o ShowSlowSqlStackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowSqlStackRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowSqlStackRequest", string(data)}, " ")
}

type ShowSlowSqlStackRequestXLanguage struct {
	value string
}

type ShowSlowSqlStackRequestXLanguageEnum struct {
	ZH_CN ShowSlowSqlStackRequestXLanguage
	EN_US ShowSlowSqlStackRequestXLanguage
}

func GetShowSlowSqlStackRequestXLanguageEnum() ShowSlowSqlStackRequestXLanguageEnum {
	return ShowSlowSqlStackRequestXLanguageEnum{
		ZH_CN: ShowSlowSqlStackRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSlowSqlStackRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSlowSqlStackRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSlowSqlStackRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSlowSqlStackRequestXLanguage) UnmarshalJSON(b []byte) error {
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
