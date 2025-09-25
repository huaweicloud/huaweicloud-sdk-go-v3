package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DockingReceptorDto 结合位点
type DockingReceptorDto struct {
	Receptor *ReceptorDrugFile `json:"receptor"`

	BoundingBox *BoundingBoxDto `json:"bounding_box"`

	// 去除受体中的离子
	RemoveIon *bool `json:"remove_ion,omitempty"`

	// 去除受体中的水分子
	RemoveWater *bool `json:"remove_water,omitempty"`

	// 去除受体中的配体分子
	RemoveLigand *bool `json:"remove_ligand,omitempty"`

	// 增加氢原子
	AddHydrogen *bool `json:"add_hydrogen,omitempty"`

	// **参数解释**： 对接类型。 **约束限制**： 不支持 **取值范围**： - BLIND_DOCKING：全局对接 - POCKET_DOCKING：口袋对接 **默认取值**： POCKET_DOCKING
	DockingType *string `json:"docking_type,omitempty"`

	ReferenceFile *ReferenceLigandFile `json:"reference_file,omitempty"`
}

func (o DockingReceptorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DockingReceptorDto struct{}"
	}

	return strings.Join([]string{"DockingReceptorDto", string(data)}, " ")
}
