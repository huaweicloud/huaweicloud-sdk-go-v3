package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoveNodesSpec struct {
	Login *Login `json:"login,omitempty"`

	// **参数解释**： 待操作节点列表，当前最多支持同时移除200个节点。 **约束限制**： 不涉及
	Nodes []NodeItem `json:"nodes"`
}

func (o RemoveNodesSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveNodesSpec struct{}"
	}

	return strings.Join([]string{"RemoveNodesSpec", string(data)}, " ")
}
