package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeApplyFailEnum **参数解释**: 节点应用成功与否状态、原因 - COLLECTOR_USE 采集器使用中，无法移除 - NODE_OFFLINE 节点失联状态，无法应用  **约束限制** 不涉及 **取值范围**: - COLLECTOR_USE - NODE_OFFLINE  **默认值** 不涉及
type NodeApplyFailEnum struct {
	value string
}

type NodeApplyFailEnumEnum struct {
	COLLECTOR_USE NodeApplyFailEnum
	NODE_OFFLINE  NodeApplyFailEnum
}

func GetNodeApplyFailEnumEnum() NodeApplyFailEnumEnum {
	return NodeApplyFailEnumEnum{
		COLLECTOR_USE: NodeApplyFailEnum{
			value: "COLLECTOR_USE",
		},
		NODE_OFFLINE: NodeApplyFailEnum{
			value: "NODE_OFFLINE",
		},
	}
}

func (c NodeApplyFailEnum) Value() string {
	return c.value
}

func (c NodeApplyFailEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeApplyFailEnum) UnmarshalJSON(b []byte) error {
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
