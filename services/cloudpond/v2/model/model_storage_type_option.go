package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageTypeOption struct {

	// 存储类型id，uuid
	Id *string `json:"id,omitempty"`

	// 售卖存储类型
	Name *string `json:"name,omitempty"`

	// 地区编码，表示允许在这个国家购买部署
	ZoneCode *string `json:"zone_code,omitempty"`

	// 步长，为0时仅可购买gears中的容量
	ExpandCapacityStep *int32 `json:"expand_capacity_step,omitempty"`

	// 固定购买容量，为空时直接按步长购买
	SupportedCapacities *[]int32 `json:"supported_capacities,omitempty"`
}

func (o StorageTypeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageTypeOption struct{}"
	}

	return strings.Join([]string{"StorageTypeOption", string(data)}, " ")
}
