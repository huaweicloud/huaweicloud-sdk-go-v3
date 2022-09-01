package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规格。
type EsflavorsVersionsFlavorsResp struct {

	// 实例的CPU核数。
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu"`

	// 实例的内存大小。单位GB。
	Ram *int32 `json:"ram,omitempty" xml:"ram"`

	// 规格名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 可用region。
	Region *string `json:"region,omitempty" xml:"region"`

	// 实例的硬盘容量范围。
	Diskrange *string `json:"diskrange,omitempty" xml:"diskrange"`

	// 可用区。
	AvailableAZ *string `json:"availableAZ,omitempty" xml:"availableAZ"`

	// 规格对应的ID。
	FlavorId *string `json:"flavor_id,omitempty" xml:"flavor_id"`
}

func (o EsflavorsVersionsFlavorsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsflavorsVersionsFlavorsResp struct{}"
	}

	return strings.Join([]string{"EsflavorsVersionsFlavorsResp", string(data)}, " ")
}
