package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TopologyNodeInfo 拓扑图节点信息
type TopologyNodeInfo struct {

	// **参数解释**: 流程图拓扑图的节点实例类型 **取值范围**: - TASK
	InstanceType *string `json:"instance_type,omitempty"`

	// **参数解释**: 流程拓扑图的节点ID **取值范围**: 不涉及
	ActionId *string `json:"action_id,omitempty"`

	// **参数解释**:     流程拓扑图的节点名称 **取值范围**: 不涉及
	ActionName *string `json:"action_name,omitempty"`

	// **参数解释**:          流程图拓扑图的节点开始时间 **取值范围**: - 不涉及
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**:          流程图拓扑图的节点结束时间 **取值范围**: - 不涉及
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**:          流程图拓扑图的节点错误信息 **取值范围**: - 不涉及
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**:          流程图拓扑图的节点输入信息 **取值范围**: - 不涉及
	Input *string `json:"input,omitempty"`

	// **参数解释**:          流程图拓扑图的节点输出信息 **取值范围**: - 不涉及
	Output *string `json:"output,omitempty"`

	// **参数解释**:          流程图拓扑图的父实例ID **取值范围**: - 不涉及
	ParentInstanceId *string `json:"parent_instance_id,omitempty"`

	// **参数解释**:          流程图拓扑图的节点的状态 **取值范围**: - RUNNING 运行中 - FAILED  运行失败 - FINISHED 运行结束
	Status *TopologyNodeInfoStatus `json:"status,omitempty"`

	// **参数解释**:          流程图拓扑图的节点是否成功 **取值范围**: - true  成功 - false 失败
	Succeed *bool `json:"succeed,omitempty"`
}

func (o TopologyNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopologyNodeInfo struct{}"
	}

	return strings.Join([]string{"TopologyNodeInfo", string(data)}, " ")
}

type TopologyNodeInfoStatus struct {
	value string
}

type TopologyNodeInfoStatusEnum struct {
	RUNNING  TopologyNodeInfoStatus
	FAILED   TopologyNodeInfoStatus
	FINISHED TopologyNodeInfoStatus
}

func GetTopologyNodeInfoStatusEnum() TopologyNodeInfoStatusEnum {
	return TopologyNodeInfoStatusEnum{
		RUNNING: TopologyNodeInfoStatus{
			value: "RUNNING",
		},
		FAILED: TopologyNodeInfoStatus{
			value: "FAILED",
		},
		FINISHED: TopologyNodeInfoStatus{
			value: "FINISHED",
		},
	}
}

func (c TopologyNodeInfoStatus) Value() string {
	return c.value
}

func (c TopologyNodeInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TopologyNodeInfoStatus) UnmarshalJSON(b []byte) error {
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
