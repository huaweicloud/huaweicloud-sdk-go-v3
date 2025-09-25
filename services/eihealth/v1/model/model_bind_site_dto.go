package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindSiteDto 结合位点
type BindSiteDto struct {

	// 靶点名称，只能设置为target1或者target2。
	Name *string `json:"name,omitempty"`

	Receptor *ReceptorDrugFile `json:"receptor"`

	// **参数解释**： 引擎[，仅支持VINA，默认值为VINA](tag:hcs)[，支持DSDP、AUTODOCK_VINA，默认值为AUTODOCK_VINA](tag:hws)。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： 对接类型[，仅支持POCKET_DOCKING](tag:hws)[，支持BLIND_DOCKING、POCKET_DOCKING](tag:hcs)。 **约束限制**： 不涉及 **取值范围**： * POCKET_DOCKING：口袋对接 * [BLIND_DOCKING：全局对接](tag:hcs) **默认取值**： POCKET_DOCKING
	DockingType *string `json:"docking_type,omitempty"`

	BoundingBox *BoundingBoxDto `json:"bounding_box,omitempty"`

	// 去除受体中的离子
	RemoveIon *bool `json:"remove_ion,omitempty"`

	// 去除受体中的水分子
	RemoveWater *bool `json:"remove_water,omitempty"`

	// 去除受体中的配体分子
	RemoveLigand *bool `json:"remove_ligand,omitempty"`

	// 增加氢原子
	AddHydrogen *bool `json:"add_hydrogen,omitempty"`
}

func (o BindSiteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindSiteDto struct{}"
	}

	return strings.Join([]string{"BindSiteDto", string(data)}, " ")
}
