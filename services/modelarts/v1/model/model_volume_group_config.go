package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeGroupConfig **参数解释**：磁盘高级配置信息。
type VolumeGroupConfig struct {

	// **参数解释**：磁盘分组名称。作为dataVolumes中volumeGroup的索引。 **取值范围**：不涉及。
	VolumeGroup string `json:"volumeGroup"`

	// **参数解释**：资源池节点容器盘占数据盘的百分比。仅磁盘分组名称为vgpaas时，即容器盘，才可指定此参数。 **取值范围**：不涉及。
	DockerThinPool *int32 `json:"dockerThinPool,omitempty"`

	LvmConfig *LvmConfig `json:"lvmConfig,omitempty"`

	// **参数解释**：存储类型。可选项如下： - volume：云硬盘。当指定dataVolumes时，该值为缺省值。 - local：本地盘。使用本地盘必须指定该字段。
	Types *[]string `json:"types,omitempty"`
}

func (o VolumeGroupConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeGroupConfig struct{}"
	}

	return strings.Join([]string{"VolumeGroupConfig", string(data)}, " ")
}
