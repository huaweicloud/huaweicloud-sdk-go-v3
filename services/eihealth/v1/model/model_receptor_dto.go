package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReceptorDto 结合位点
type ReceptorDto struct {

	// 靶点名称。
	Name *string `json:"name,omitempty"`

	Receptor *ReceptorDrugFile `json:"receptor"`

	BoundingBox *BoundBoxDto `json:"bounding_box,omitempty"`

	// 去除受体中的离子。
	RemoveIon *bool `json:"remove_ion,omitempty"`

	// 去除受体中的水分子。
	RemoveWater *bool `json:"remove_water,omitempty"`

	// 去除受体中的配体分子。
	RemoveLigand *bool `json:"remove_ligand,omitempty"`

	// 增加氢原子。
	AddHydrogen *bool `json:"add_hydrogen,omitempty"`
}

func (o ReceptorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReceptorDto struct{}"
	}

	return strings.Join([]string{"ReceptorDto", string(data)}, " ")
}
