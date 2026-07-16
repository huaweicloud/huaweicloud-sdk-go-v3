package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatusClusters **参数解释**：资源池集群信息，特权池才有该字段。
type PoolStatusClusters struct {

	// **参数解释**：集群名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：标准池的集群规格。 **取值范围**：不涉及。
	ClusterFlavor *string `json:"clusterFlavor,omitempty"`

	// **参数解释**：资源池集群的类型。 **取值范围**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：集群的版本号。 **取值范围**：不涉及。
	Version *string `json:"version,omitempty"`

	Plugins *PoolStatusClustersPlugins `json:"plugins,omitempty"`
}

func (o PoolStatusClusters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatusClusters struct{}"
	}

	return strings.Join([]string{"PoolStatusClusters", string(data)}, " ")
}
