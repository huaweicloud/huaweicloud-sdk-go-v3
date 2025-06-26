package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskResp struct {

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 主机名称。 **取值范围**： 不涉及。
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 磁盘名称。 **取值范围**： 不涉及。
	DiskName *string `json:"disk_name,omitempty"`

	// **参数解释**： 磁盘类型(系统盘、数据盘、日志盘)。 **取值范围**： 不涉及。
	DiskType *string `json:"disk_type,omitempty"`

	// **参数解释**： 磁盘总容量(GB)。 **取值范围**： 不涉及。
	Total *float64 `json:"total,omitempty"`

	// **参数解释**： 磁盘已使用容量(GB)。 **取值范围**： 不涉及。
	Used *float64 `json:"used,omitempty"`

	// **参数解释**： 磁盘可用容量(GB)。 **取值范围**： 不涉及。
	Available *float64 `json:"available,omitempty"`

	// **参数解释**： 磁盘使用率(%)。 **取值范围**： 不涉及。
	UsedPercentage *float64 `json:"used_percentage,omitempty"`

	// **参数解释**： IO等待时间(ms)。 **取值范围**： 不涉及。
	Await *float64 `json:"await,omitempty"`

	// **参数解释**： IO服务时间(ms)。 **取值范围**： 不涉及。
	Svctm *float64 `json:"svctm,omitempty"`

	// **参数解释**： IO使用率(%)。 **取值范围**： 不涉及。
	Util *float64 `json:"util,omitempty"`

	// **参数解释**： 磁盘读速率(KB/s)。 **取值范围**： 不涉及。
	ReadRate *float64 `json:"read_rate,omitempty"`

	// **参数解释**： 磁盘写速率(KB/s)。 **取值范围**： 不涉及。
	WriteRate *float64 `json:"write_rate,omitempty"`
}

func (o DiskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskResp struct{}"
	}

	return strings.Join([]string{"DiskResp", string(data)}, " ")
}
