package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigUpdateSpecConfigs **参数解释**：待更新的节点配置。支持镜像预热（image_warm_up）、CPU管理策略（cpu_manager）、缓存清理（drop_cache）、透明大页（transparent_hugepage）等配置项。 **约束限制**：仅Dedicate（专属）资源池支持修改节点配置，非专属资源池调用将返回403。镜像预热配置仅在资源池机型为支持镜像预热的裸金属规格（Ascend 910B/910C）时可用。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigUpdateSpecConfigs struct {

	// **参数解释**：镜像预热配置。用于创建或删除镜像预热任务，镜像会被提前拉取到资源池节点，缩短训练任务启动时间。 **约束限制**：镜像预热任务数量上限为10；同一请求内任务名称不可重复、镜像地址不可重复；仅标准池支持，Lite资源池不支持该配置。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ImageWarmUp *[]NodeconfigUpdateSpecConfigsImageWarmUp `json:"image_warm_up,omitempty"`

	CpuManager *NodeconfigUpdateSpecConfigsCpuManager `json:"cpu_manager,omitempty"`

	DropCache *NodeconfigUpdateSpecConfigsDropCache `json:"drop_cache,omitempty"`

	TransparentHugepage *NodeconfigUpdateSpecConfigsTransparentHugepage `json:"transparent_hugepage,omitempty"`
}

func (o NodeconfigUpdateSpecConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigUpdateSpecConfigs struct{}"
	}

	return strings.Join([]string{"NodeconfigUpdateSpecConfigs", string(data)}, " ")
}
