package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkDeviceSpec 网络设备规格
type NetworkDeviceSpec struct {

	// 网络设备规格ID
	Id *string `json:"id,omitempty"`

	// 网络设备规格名称。
	Name *string `json:"name,omitempty"`

	// 设备功率。单位：w
	Power *int32 `json:"power,omitempty"`

	// 设备高度。U位数。
	Unit *int32 `json:"unit,omitempty"`

	// 用途
	PerformanceType *string `json:"performance_type,omitempty"`
}

func (o NetworkDeviceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkDeviceSpec struct{}"
	}

	return strings.Join([]string{"NetworkDeviceSpec", string(data)}, " ")
}
