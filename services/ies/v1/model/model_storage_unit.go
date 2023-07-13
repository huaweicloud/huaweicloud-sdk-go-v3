package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageUnit 存储单元
type StorageUnit struct {
	StorageType *StorageType `json:"storage_type"`

	// 存储池大小，单位：TB。
	Capacity int32 `json:"capacity"`

	// 存储池销售档位
	Gears []int32 `json:"gears"`

	// 规格类型。例如：highio-4T
	FlavorType string `json:"flavor_type"`

	// 存储节点台数。
	Count int32 `json:"count"`
}

func (o StorageUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageUnit struct{}"
	}

	return strings.Join([]string{"StorageUnit", string(data)}, " ")
}
