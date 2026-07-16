package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModelVolumeGroupConfigs **参数解释**：磁盘高级配置信息。
type PoolSpecModelVolumeGroupConfigs struct {

	// **参数解释**：磁盘分组名称。作为dataVolumes中volumeGroup的索引。 **取值范围**：不涉及。
	VolumeGroup *string `json:"volumeGroup,omitempty"`

	// **参数解释**：资源池节点容器盘占数据盘的百分比。仅磁盘分组名称为vgpaas时，即容器盘，才可指定此参数。 **取值范围**：不涉及。
	DockerThinPool *int32 `json:"dockerThinPool,omitempty"`

	LvmConfig *PoolSpecModelVolumeGroupConfigsLvmConfig `json:"lvmConfig,omitempty"`
}

func (o PoolSpecModelVolumeGroupConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelVolumeGroupConfigs struct{}"
	}

	return strings.Join([]string{"PoolSpecModelVolumeGroupConfigs", string(data)}, " ")
}
