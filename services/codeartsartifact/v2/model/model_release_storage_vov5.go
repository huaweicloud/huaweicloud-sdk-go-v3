package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReleaseStorageVov5 struct {

	// **参数解释**: 已使用存储量 ---带单位。 **取值范围**: 不涉及。
	UsedCapacity *string `json:"used_capacity,omitempty"`

	// **参数解释**: 租户存储最大量 ---带单位。 **取值范围**: 不涉及。
	TotalCapacity *string `json:"total_capacity,omitempty"`

	// **参数解释**: 已使用存储量 ---字节。 **取值范围**: 不涉及。
	UsedCapacitySize *int64 `json:"used_capacity_size,omitempty"`

	// **参数解释**: 租户存储最大量 ---字节。 **取值范围**: 不涉及。
	TotalCapacitySize *int64 `json:"total_capacity_size,omitempty"`

	// **参数解释**: 包周期套餐类型(0.no_package 1.basic 2.professional 3.platinum)。 **取值范围**: 不涉及。
	PackageType *string `json:"package_type,omitempty"`

	// **参数解释**: 项目仓库数量。 **取值范围**: 不涉及。
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (o ReleaseStorageVov5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReleaseStorageVov5 struct{}"
	}

	return strings.Join([]string{"ReleaseStorageVov5", string(data)}, " ")
}
