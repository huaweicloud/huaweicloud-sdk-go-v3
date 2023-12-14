package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeConstraints 节点限制
type NodeConstraints struct {

	// 其他节点限制
	Other map[string]interface{} `json:"other,omitempty"`

	Master *NodeConstraint `json:"master,omitempty"`

	Core *NodeConstraint `json:"core,omitempty"`

	Task *NodeConstraint `json:"task,omitempty"`

	CoreSeparate *NodeConstraint `json:"core_separate,omitempty"`

	CoreCombine *NodeConstraint `json:"core_combine,omitempty"`

	TaskSeparate *NodeConstraint `json:"task_separate,omitempty"`

	TaskCombine *NodeConstraint `json:"task_combine,omitempty"`

	NodeGroupTask *NodeConstraint `json:"node_group_task,omitempty"`
}

func (o NodeConstraints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeConstraints struct{}"
	}

	return strings.Join([]string{"NodeConstraints", string(data)}, " ")
}
