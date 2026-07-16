package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RootVolume 系统盘信息。
type RootVolume struct {

	// **参数解释**：磁盘类型[，具体内容可参考磁盘类型及性能介绍](tag:hc)。 **取值范围**：可选值如下： - SSD：超高IO硬盘 - GPSSD：通用型SSD - SAS：高IO硬盘
	VolumeType string `json:"volumeType"`

	// **参数解释**：磁盘大小，单位为GiB。 **取值范围**：不涉及。
	Size string `json:"size"`
}

func (o RootVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RootVolume struct{}"
	}

	return strings.Join([]string{"RootVolume", string(data)}, " ")
}
