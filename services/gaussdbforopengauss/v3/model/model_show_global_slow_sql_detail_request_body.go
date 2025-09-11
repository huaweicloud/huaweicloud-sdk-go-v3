package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowGlobalSlowSqlDetailRequestBody struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 节点ID列表。 **约束限制**: 节点ID数组不能为空。
	NodeIds *[]string `json:"node_ids,omitempty"`

	// **参数解释**: 起始日期。 **约束限制**: 13位UNIX时间戳格式，单位是毫秒，时区是UTC。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**: 结束日期。 **约束限制**: 13位UNIX时间戳格式，单位是毫秒，时区是UTC。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**: 慢SQL的SQL ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId string `json:"sql_id"`

	// **参数解释**: 组件类型。 **约束限制**: 不涉及。 **取值范围**: - cn：CN组件类型。 - dn：DN组件类型。  **默认取值**: 不涉及。
	ComponentType ShowGlobalSlowSqlDetailRequestBodyComponentType `json:"component_type"`
}

func (o ShowGlobalSlowSqlDetailRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalSlowSqlDetailRequestBody struct{}"
	}

	return strings.Join([]string{"ShowGlobalSlowSqlDetailRequestBody", string(data)}, " ")
}

type ShowGlobalSlowSqlDetailRequestBodyComponentType struct {
	value string
}

type ShowGlobalSlowSqlDetailRequestBodyComponentTypeEnum struct {
	CN ShowGlobalSlowSqlDetailRequestBodyComponentType
	DN ShowGlobalSlowSqlDetailRequestBodyComponentType
}

func GetShowGlobalSlowSqlDetailRequestBodyComponentTypeEnum() ShowGlobalSlowSqlDetailRequestBodyComponentTypeEnum {
	return ShowGlobalSlowSqlDetailRequestBodyComponentTypeEnum{
		CN: ShowGlobalSlowSqlDetailRequestBodyComponentType{
			value: "cn",
		},
		DN: ShowGlobalSlowSqlDetailRequestBodyComponentType{
			value: "dn",
		},
	}
}

func (c ShowGlobalSlowSqlDetailRequestBodyComponentType) Value() string {
	return c.value
}

func (c ShowGlobalSlowSqlDetailRequestBodyComponentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGlobalSlowSqlDetailRequestBodyComponentType) UnmarshalJSON(b []byte) error {
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
