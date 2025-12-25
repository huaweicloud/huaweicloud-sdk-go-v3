package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeVo struct {

	// 相关描述信息
	Args *[]ArgsVo `json:"args,omitempty"`

	// UUID
	NodeId *string `json:"node_id,omitempty"`

	NodeStatus *NodeStatus `json:"node_status,omitempty"`
}

func (o NodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeVo struct{}"
	}

	return strings.Join([]string{"NodeVo", string(data)}, " ")
}
