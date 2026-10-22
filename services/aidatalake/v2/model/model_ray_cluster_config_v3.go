package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayClusterConfigV3 **参数解释**：Ray集群规格。
type RayClusterConfigV3 struct {
	HeadGroupSpec *HeadGroupSpecV3 `json:"head_group_spec"`

	// **参数解释**：Worker Group的配置。 **约束限制**：不涉及。
	WorkerGroupSpecs []WorkerGroupSpecV3 `json:"worker_group_specs"`
}

func (o RayClusterConfigV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayClusterConfigV3 struct{}"
	}

	return strings.Join([]string{"RayClusterConfigV3", string(data)}, " ")
}
