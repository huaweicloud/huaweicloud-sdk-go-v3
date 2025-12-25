package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeMessage struct {

	// 角色
	Role *string `json:"role,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 原始内容
	Origin *string `json:"origin,omitempty"`

	// 节点id
	NodeId *string `json:"nodeId,omitempty"`

	// 节点类型
	NodeType *string `json:"nodeType,omitempty"`

	// 节点名
	NodeName *string `json:"nodeName,omitempty"`

	// 事件发生时间
	CreatedTime *int64 `json:"createdTime,omitempty"`
}

func (o NodeMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeMessage struct{}"
	}

	return strings.Join([]string{"NodeMessage", string(data)}, " ")
}
