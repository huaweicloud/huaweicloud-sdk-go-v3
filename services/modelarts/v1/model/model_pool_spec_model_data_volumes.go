package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModelDataVolumes **参数解释**：自定义数据盘（云硬盘）列表信息，指定后不可修改。
type PoolSpecModelDataVolumes struct {

	// **参数解释**：磁盘类型，具体内容可参考磁盘类型及性能介绍。 **取值范围**：可选值如下： - SSD：超高IO硬盘 - GPSSD：通用型SSD - SAS：高IO硬盘
	VolumeType *string `json:"volumeType,omitempty"`

	// **参数解释**：磁盘大小，单位为Gi。 **取值范围**：不涉及。
	Size *string `json:"size,omitempty"`

	// **参数解释**：磁盘个数。 **取值范围**：不涉及。
	Count *string `json:"count,omitempty"`

	ExtendParams *PoolSpecModelDataVolumesExtendParams `json:"extendParams,omitempty"`
}

func (o PoolSpecModelDataVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelDataVolumes struct{}"
	}

	return strings.Join([]string{"PoolSpecModelDataVolumes", string(data)}, " ")
}
