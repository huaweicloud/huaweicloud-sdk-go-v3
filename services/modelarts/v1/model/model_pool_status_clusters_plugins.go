package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatusClustersPlugins **参数解释**：集群支持的插件名称。
type PoolStatusClustersPlugins struct {

	// **参数解释**：集群支持的插件名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o PoolStatusClustersPlugins) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatusClustersPlugins struct{}"
	}

	return strings.Join([]string{"PoolStatusClustersPlugins", string(data)}, " ")
}
