package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RingHost **参数解释**： 集群主机信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RingHost struct {

	// **参数解释**： 主机名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HostName string `json:"host_name"`

	// **参数解释**： 后端IP地址。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BackIp string `json:"back_ip"`

	// **参数解释**： 主机CPU核数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CpuCores int32 `json:"cpu_cores"`

	// **参数解释**： 主机内存。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Memory float64 `json:"memory"`

	// **参数解释**： 主机磁盘大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DiskSize float64 `json:"disk_size"`
}

func (o RingHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RingHost struct{}"
	}

	return strings.Join([]string{"RingHost", string(data)}, " ")
}
