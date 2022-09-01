package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云硬盘类型的规格
type VolumeTypeExtraSpecs struct {

	// 支持当前云硬盘类型的可用区列表，列表的元素以逗号分隔。
	RESKEYavailabilityZones *string `json:"RESKEY:availability_zones,omitempty" xml:"RESKEY:availability_zones"`

	// 预留属性。
	AvailabilityZone *string `json:"availability-zone,omitempty" xml:"availability-zone"`

	// 当前云硬盘类型已售罄的可用区列表，列表的元素以逗号分隔。
	OsVendorExtendedsoldOutAvailabilityZones *string `json:"os-vendor-extended:sold_out_availability_zones,omitempty" xml:"os-vendor-extended:sold_out_availability_zones"`

	// 预留属性。
	VolumeBackendName *string `json:"volume_backend_name,omitempty" xml:"volume_backend_name"`

	// 预留属性。
	HWavailabilityZone *string `json:"HW:availability_zone,omitempty" xml:"HW:availability_zone"`
}

func (o VolumeTypeExtraSpecs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTypeExtraSpecs struct{}"
	}

	return strings.Join([]string{"VolumeTypeExtraSpecs", string(data)}, " ")
}
