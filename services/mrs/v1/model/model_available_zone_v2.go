package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可用区信息
type AvailableZoneV2 struct {

	// 可用区编码
	Id *string `json:"id,omitempty"`

	// 可用区编码
	AzCode *string `json:"az_code,omitempty"`

	// 可用区名称
	AzName *string `json:"az_name,omitempty"`

	// 可用区id
	AzId *string `json:"az_id,omitempty"`

	// 可用区状态
	Status *string `json:"status,omitempty"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 可用区分组id
	AzGroupId *string `json:"az_group_id,omitempty"`

	// 当前AZ的类型 Core 核心 Satellite 卫星 Dedicated 专属 Virtual 虚拟 Edge 边缘 EdgeCental 中心边缘 Hybrid 混合云
	AzType *string `json:"az_type,omitempty"`

	AzTags *AvailableTag `json:"az_tags,omitempty"`
}

func (o AvailableZoneV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZoneV2 struct{}"
	}

	return strings.Join([]string{"AvailableZoneV2", string(data)}, " ")
}
