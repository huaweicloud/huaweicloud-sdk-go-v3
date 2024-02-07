package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PocketMolDesignReceptorDto 靶点口袋分子设计的靶点文件
type PocketMolDesignReceptorDto struct {
	File *ReceptorDrugFile `json:"file"`

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

func (o PocketMolDesignReceptorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PocketMolDesignReceptorDto struct{}"
	}

	return strings.Join([]string{"PocketMolDesignReceptorDto", string(data)}, " ")
}
