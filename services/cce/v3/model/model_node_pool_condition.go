package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodePoolCondition 节点池/伸缩组详细状态。
type NodePoolCondition struct {

	// **参数解释**： 状态类型。 **约束限制**： 不涉及 **取值范围**： - \"TaintSynchronizing\": 节点池正在同步节点K8s污点，不影响节点池可扩容状态（该状态类型为节点池级别，伸缩组中无该状态类型）。 - \"LabelSynchronizing\": 节点池正在同步节点K8s标签，不影响节点池可扩容状态（该状态类型为节点池级别，伸缩组中无该状态类型）。 - \"UserTagsSynchronizing\": 节点池正在同步节点资源标签，不影响节点池可扩容状态（该状态类型为节点池级别，伸缩组中无该状态类型）。 - \"ConfigurationSynchronizing\": 节点池正在同步节点配置，不影响节点池可扩容状态（该状态类型为节点池级别，伸缩组中无该状态类型）。 - \"Scalable\"：节点池/伸缩组实际的可扩容状态，如果状态为\"False\"时则不会再次触发节点池扩容行为。 - \"QuotaInsufficient\"：节点池/伸缩组扩容依赖的配额不足，影响节点池可扩容状态。 - \"ResourceInsufficient\"：节点池/伸缩组扩容依赖的资源不足，影响节点池可扩容状态。 - \"UnexpectedError\"：节点池/伸缩组非预期扩容失败，影响节点池可扩容状态。 [- \"LockedByOrder\"：节点池/伸缩组被订单锁定，此时Reason为待支付订单ID。](tag:hws,hws_hk) - \"Error\"：节点池/伸缩组错误，通常由于删除失败触发。  **默认取值**： 不涉及
	Type *NodePoolConditionType `json:"type,omitempty"`

	// Condition当前状态，取值如下 - \"True\" - \"False\"
	Status *string `json:"status,omitempty"`

	// 上次状态检查时间。
	LastProbeTime *string `json:"lastProbeTime,omitempty"`

	// 上次状态变更时间。
	LastTransitTime *string `json:"lastTransitTime,omitempty"`

	// 上次状态变更原因。
	Reason *string `json:"reason,omitempty"`

	// Condition详细描述。
	Message *string `json:"message,omitempty"`
}

func (o NodePoolCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolCondition struct{}"
	}

	return strings.Join([]string{"NodePoolCondition", string(data)}, " ")
}

type NodePoolConditionType struct {
	value string
}

type NodePoolConditionTypeEnum struct {
	TAINT_SYNCHRONIZING         NodePoolConditionType
	LABEL_SYNCHRONIZING         NodePoolConditionType
	USER_TAGS_SYNCHRONIZING     NodePoolConditionType
	CONFIGURATION_SYNCHRONIZING NodePoolConditionType
	SCALABLE                    NodePoolConditionType
	QUOTA_INSUFFICIENT          NodePoolConditionType
	RESOURCE_INSUFFICIENT       NodePoolConditionType
	UNEXPECTED_ERROR            NodePoolConditionType
	LOCKED_BY_ORDER             NodePoolConditionType
	ERROR                       NodePoolConditionType
}

func GetNodePoolConditionTypeEnum() NodePoolConditionTypeEnum {
	return NodePoolConditionTypeEnum{
		TAINT_SYNCHRONIZING: NodePoolConditionType{
			value: "TaintSynchronizing",
		},
		LABEL_SYNCHRONIZING: NodePoolConditionType{
			value: "LabelSynchronizing",
		},
		USER_TAGS_SYNCHRONIZING: NodePoolConditionType{
			value: "UserTagsSynchronizing",
		},
		CONFIGURATION_SYNCHRONIZING: NodePoolConditionType{
			value: "ConfigurationSynchronizing",
		},
		SCALABLE: NodePoolConditionType{
			value: "Scalable",
		},
		QUOTA_INSUFFICIENT: NodePoolConditionType{
			value: "QuotaInsufficient",
		},
		RESOURCE_INSUFFICIENT: NodePoolConditionType{
			value: "ResourceInsufficient",
		},
		UNEXPECTED_ERROR: NodePoolConditionType{
			value: "UnexpectedError",
		},
		LOCKED_BY_ORDER: NodePoolConditionType{
			value: "LockedByOrder",
		},
		ERROR: NodePoolConditionType{
			value: "Error",
		},
	}
}

func (c NodePoolConditionType) Value() string {
	return c.value
}

func (c NodePoolConditionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodePoolConditionType) UnmarshalJSON(b []byte) error {
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
