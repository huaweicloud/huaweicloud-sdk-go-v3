package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DockingReceptorDto 结合位点
type DockingReceptorDto struct {
	Receptor *DrugFile `json:"receptor"`

	BoundingBox *BoundingBoxDto `json:"bounding_box"`

	// 去除受体中的离子
	RemoveIon *bool `json:"remove_ion,omitempty"`

	// 去除受体中的水分子
	RemoveWater *bool `json:"remove_water,omitempty"`

	// 去除受体中的配体分子
	RemoveLigand *bool `json:"remove_ligand,omitempty"`
}

func (o DockingReceptorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DockingReceptorDto struct{}"
	}

	return strings.Join([]string{"DockingReceptorDto", string(data)}, " ")
}
