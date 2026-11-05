package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigUpdateSpec **参数解释**：待更新的节点配置插件实例请求信息。 **约束限制**：不涉及。
type NodeconfigUpdateSpec struct {
	Configs *NodeconfigUpdateSpecConfigs `json:"configs,omitempty"`
}

func (o NodeconfigUpdateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdateSpec struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdateSpec", string(data)}, " ")
}
