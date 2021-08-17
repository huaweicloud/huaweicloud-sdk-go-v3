package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type NodePoolStatus struct {
	// 当前节点池中所有节点数量（不含删除中的节点）。

	CurrentNode *int32 `json:"currentNode,omitempty"`
	// 当前节点池中处于创建流程中的节点数量。

	CreatingNode *int32 `json:"creatingNode,omitempty"`
	// 当前节点池中删除中或者删除失败的节点数量。

	DeletingNode *int32 `json:"deletingNode,omitempty"`
	// 节点池状态，为空时节点池处于可用状态。 - Synchronizing：伸缩中 - Synchronized：节点池更新失败时会被置于此状态 - SoldOut：节点资源售罄 - Deleting：删除中 - Error：错误

	Phase *NodePoolStatusPhase `json:"phase,omitempty"`
	// 对节点池执行操作时的 JobID。

	JobId *string `json:"jobId,omitempty"`
	// 节点池每次扩容的动作结果记录，用于确定节点池是否还能继续扩容。

	Conditions *[]NodePoolCondition `json:"conditions,omitempty"`
}

func (o NodePoolStatus) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodePoolStatus struct{}"
	}

	return strings.Join([]string{"NodePoolStatus", string(data)}, " ")
}

type NodePoolStatusPhase struct {
	value string
}

type NodePoolStatusPhaseEnum struct {
	SYNCHRONIZING NodePoolStatusPhase
	SYNCHRONIZED  NodePoolStatusPhase
	SOLD_OUT      NodePoolStatusPhase
	DELETING      NodePoolStatusPhase
	ERROR         NodePoolStatusPhase
}

func GetNodePoolStatusPhaseEnum() NodePoolStatusPhaseEnum {
	return NodePoolStatusPhaseEnum{
		SYNCHRONIZING: NodePoolStatusPhase{
			value: "Synchronizing",
		},
		SYNCHRONIZED: NodePoolStatusPhase{
			value: "Synchronized",
		},
		SOLD_OUT: NodePoolStatusPhase{
			value: "SoldOut",
		},
		DELETING: NodePoolStatusPhase{
			value: "Deleting",
		},
		ERROR: NodePoolStatusPhase{
			value: "Error",
		},
	}
}

func (c NodePoolStatusPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *NodePoolStatusPhase) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
