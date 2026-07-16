package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModelVolumeGroupConfigsLvmConfig **参数解释**：LVM配置管理。
type PoolSpecModelVolumeGroupConfigsLvmConfig struct {

	// **参数解释**：LVM写入模式 **取值范围**：可选项如下： - linear：线性模式。 - striped：条带模式，使用多块磁盘组成条带模式，能够提升磁盘性能。
	LvType *string `json:"lvType,omitempty"`
}

func (o PoolSpecModelVolumeGroupConfigsLvmConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelVolumeGroupConfigsLvmConfig struct{}"
	}

	return strings.Join([]string{"PoolSpecModelVolumeGroupConfigsLvmConfig", string(data)}, " ")
}
