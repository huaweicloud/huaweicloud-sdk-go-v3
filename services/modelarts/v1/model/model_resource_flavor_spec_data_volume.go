package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceFlavorSpecDataVolume struct {

	// **参数解释**：磁盘类型[，具体内容可参考[磁盘类型及性能介绍](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)](tag:hc)。 **取值范围**：可选值如下： - SSD：超高IO硬盘 - GPSSD：通用型SSD - SAS：高IO硬盘
	VolumeType *string `json:"volumeType,omitempty"`

	// **参数解释**：磁盘大小，单位为Gi。 **取值范围**：不涉及。
	Size *string `json:"size,omitempty"`
}

func (o ResourceFlavorSpecDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorSpecDataVolume struct{}"
	}

	return strings.Join([]string{"ResourceFlavorSpecDataVolume", string(data)}, " ")
}
