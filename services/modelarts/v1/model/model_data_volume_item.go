package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataVolumeItem 数据盘信息。
type DataVolumeItem struct {

	// **参数解释**：磁盘类型[，具体内容可参考磁盘类型及性能介绍](tag:hc)。 **取值范围**：可选值如下： - SSD：超高IO硬盘 - GPSSD：通用型SSD - SAS：高IO硬盘
	VolumeType string `json:"volumeType"`

	// **参数解释**：磁盘大小，单位为GiB。 **取值范围**：不涉及。
	Size string `json:"size"`

	// **参数解释**：磁盘个数。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	ExtendParams *VolumeExtendParams `json:"extendParams"`
}

func (o DataVolumeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataVolumeItem struct{}"
	}

	return strings.Join([]string{"DataVolumeItem", string(data)}, " ")
}
