package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Node 任务节点定义。
type Node struct {

	// 父亲节点的名称。
	ParentNode *string `json:"parent_node,omitempty"`

	// 节点类型。
	Category *string `json:"category,omitempty"`

	// 节点描述。
	Description *string `json:"description,omitempty"`

	// 节点id
	Id *string `json:"id,omitempty"`

	// 是否忽略错误
	IgnoreError *bool `json:"ignore_error,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	// 节点名称,比如是Node。
	Name *string `json:"name,omitempty"`

	// 任务名称，节点上任务的名称。
	TaskName *string `json:"task_name,omitempty"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
