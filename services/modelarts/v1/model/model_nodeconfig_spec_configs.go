package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigSpecConfigs **参数解释**： 节点自定义配置；当前支持节点绑核（cpu_manager）、缓存清理（drop_cache）、透明大页（transparent_hugepage）和镜像预热（image_warm_up）。 **约束限制**：仅Dedicate（专属）资源池支持修改节点配置，非专属资源池调用将返回403。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigSpecConfigs struct {

	// **参数解释**：镜像预热任务配置列表。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ImageWarmUp *[]NodeconfigSpecConfigsImageWarmUp `json:"image_warm_up,omitempty"`

	CpuManager *NodeconfigSpecConfigsCpuManager `json:"cpu_manager,omitempty"`

	DropCache *NodeconfigSpecConfigsDropCache `json:"drop_cache,omitempty"`

	TransparentHugepage *NodeconfigSpecConfigsTransparentHugepage `json:"transparent_hugepage,omitempty"`
}

func (o NodeconfigSpecConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigSpecConfigs struct{}"
	}

	return strings.Join([]string{"NodeconfigSpecConfigs", string(data)}, " ")
}
