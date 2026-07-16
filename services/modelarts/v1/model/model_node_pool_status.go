package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolStatus 节点池状态信息。
type NodePoolStatus struct {
	Resources *NodePoolStatusResources `json:"resources,omitempty"`
}

func (o NodePoolStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolStatus struct{}"
	}

	return strings.Join([]string{"NodePoolStatus", string(data)}, " ")
}
