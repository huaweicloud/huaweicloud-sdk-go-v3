package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListProductsRespIo struct {

	// IO类型。
	IoType *string `json:"io_type,omitempty" xml:"io_type"`

	// IO规格。
	StorageSpecCode *string `json:"storage_spec_code,omitempty" xml:"storage_spec_code"`

	// IO未售罄的可用区列表。
	AvailableZones *[]string `json:"available_zones,omitempty" xml:"available_zones"`

	// IO已售罄的不可用区列表。
	UnavailableZones *[]string `json:"unavailable_zones,omitempty" xml:"unavailable_zones"`

	// 磁盘类型。
	VolumeType *string `json:"volume_type,omitempty" xml:"volume_type"`
}

func (o ListProductsRespIo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductsRespIo struct{}"
	}

	return strings.Join([]string{"ListProductsRespIo", string(data)}, " ")
}
