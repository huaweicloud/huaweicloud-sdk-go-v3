package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Flavor struct {
	// 边缘实例规格的ID。

	Id *string `json:"id,omitempty"`
	// 边缘实例规格的名称。

	Name *string `json:"name,omitempty"`
	// 边缘实例规格对应要求系统盘大小。  当前未使用该参数，缺省值为0。

	Disk *string `json:"disk,omitempty"`
	// 边缘实例规格对应的内存大小，单位为MB。

	Ram *int32 `json:"ram,omitempty"`
	// 边缘实例规格对应的CPU核数。

	Vcpus *string `json:"vcpus,omitempty"`

	OsExtraSpecs *OsExtraSpecs `json:"os_extra_specs,omitempty"`
	// 扩展属性，flavor是否给所有租户使用。

	OsFlavorAccessIsPublic *bool `json:"os_flavor_access_is_public,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
