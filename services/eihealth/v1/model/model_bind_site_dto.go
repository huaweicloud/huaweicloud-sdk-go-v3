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
