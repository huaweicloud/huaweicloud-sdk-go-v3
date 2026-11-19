package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorCategoryDetail struct {

	// **参数解释**: 计算规格组列表。 **取值范围**: 不涉及。
	ComputeFlavorGroups *[]ComputeFlavorGroupDetail `json:"compute_flavor_groups,omitempty"`

	// **参数解释**: 存储规格列表。 **取值范围**: 不涉及。
	VolumeFlavors *[]VolumeFlavorItemDetail `json:"volume_flavors,omitempty"`

	// **参数解释**: DeC专属规格列表。 **取值范围**: 不涉及。
	DecServiceFlavors *[]DecServiceFlavorItemDetail `json:"dec_service_flavors,omitempty"`
}

func (o FlavorCategoryDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorCategoryDetail struct{}"
	}

	return strings.Join([]string{"FlavorCategoryDetail", string(data)}, " ")
}
