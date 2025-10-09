package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterSpec master的配置，支持指定可用区、规格和故障域。若指定故障域，则必须所有master节点都需要指定故障字段。
type MasterSpec struct {

	// **参数解释：** 控制节点所在的可用区，需要指定可用区（AZ）的名称。 [CCE支持的可用区请参考[地区和终端节点](https://console.huaweicloud.com/apiexplorer/#/endpoint/CCE)](tag:hws) [CCE支持的可用区请参考[地区和终端节点](https://console-intl.huaweicloud.com/apiexplorer/#/endpoint/CCE)](tag:hws_hk) **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// 规格
	Flavor *string `json:"flavor,omitempty"`

	// 故障域。 1. 指定该字段需要当前系统已开启故障域特性，否则校验失败。 2. 仅单az场景支持且必须显式指定az。
	FaultDomain *string `json:"faultDomain,omitempty"`
}

func (o MasterSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterSpec struct{}"
	}

	return strings.Join([]string{"MasterSpec", string(data)}, " ")
}
