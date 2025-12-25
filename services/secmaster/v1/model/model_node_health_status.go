package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeHealthStatus **参数解释**: 节点的健康状态 - NORMAL 正常 - ANOMALIES 异常 - FAULTS 故障 - LOST_CONTACT 失联  **约束限制** 不涉及 **取值范围**: - NORMAL - ANOMALIES - FAULTS - LOST_CONTACT  **默认值** 不涉及
type NodeHealthStatus struct {
	value string
}

type NodeHealthStatusEnum struct {
	NORMAL       NodeHealthStatus
	ANOMALIES    NodeHealthStatus
	FAULTS       NodeHealthStatus
	LOST_CONTACT NodeHealthStatus
}

func GetNodeHealthStatusEnum() NodeHealthStatusEnum {
	return NodeHealthStatusEnum{
		NORMAL: NodeHealthStatus{
			value: "NORMAL",
		},
		ANOMALIES: NodeHealthStatus{
			value: "ANOMALIES",
		},
		FAULTS: NodeHealthStatus{
			value: "FAULTS",
		},
		LOST_CONTACT: NodeHealthStatus{
			value: "LOST_CONTACT",
		},
	}
}

func (c NodeHealthStatus) Value() string {
	return c.value
}

func (c NodeHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeHealthStatus) UnmarshalJSON(b []byte) error {
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
