package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeVo **参数解释**：磁盘信息。 **约束限制**：不涉及。
type VolumeVo struct {

	// **参数解释**：磁盘类型[，具体内容可参考磁盘类型及性能介绍](tag:hc)。 **取值范围**：   - SSD：超高IO硬盘   - GPSSD：通用型SSD
	VolumeType string `json:"volumeType"`

	// **参数解释**：磁盘大小，单位为Gi。 **取值范围**：不涉及。
	Size string `json:"size"`

	// **参数解释**：磁盘个数。不指定默认值为1。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`
}

func (o VolumeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeVo struct{}"
	}

	return strings.Join([]string{"VolumeVo", string(data)}, " ")
}
