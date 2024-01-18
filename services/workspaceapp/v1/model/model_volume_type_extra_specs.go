package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeTypeExtraSpecs struct {

	// 该类型磁盘对应的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 已售罄的磁盘可用区。
	SoldOutAvailabilityZone *string `json:"sold_out_availability_zone,omitempty"`
}

func (o VolumeTypeExtraSpecs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTypeExtraSpecs struct{}"
	}

	return strings.Join([]string{"VolumeTypeExtraSpecs", string(data)}, " ")
}
