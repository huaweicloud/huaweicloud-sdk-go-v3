package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeStatus **参数解释**: 节点状态 - RUN 运行 - STOP 停止  **约束限制** 不涉及 **取值范围**: - RUN - STOP  **默认值** 不涉及
type NodeStatus struct {
	value string
}

type NodeStatusEnum struct {
	RUN  NodeStatus
	STOP NodeStatus
}

func GetNodeStatusEnum() NodeStatusEnum {
	return NodeStatusEnum{
		RUN: NodeStatus{
			value: "RUN",
		},
		STOP: NodeStatus{
			value: "STOP",
		},
	}
}

func (c NodeStatus) Value() string {
	return c.value
}

func (c NodeStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeStatus) UnmarshalJSON(b []byte) error {
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
