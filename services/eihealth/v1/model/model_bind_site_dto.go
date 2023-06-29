package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindSiteDto 结合位点
type BindSiteDto struct {
	Receptor *DrugFile `json:"receptor"`

	BoundingBox *BoundingBoxDto `json:"bounding_box,omitempty"`

	// 去除受体中的离子
	RemoveIon *bool `json:"remove_ion,omitempty"`

	// 去除受体中的水分子
	RemoveWater *bool `json:"remove_water,omitempty"`

	// 去除受体中的配体分子
	RemoveLigand *bool `json:"remove_ligand,omitempty"`
}

func (o BindSiteDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindSiteDto struct{}"
	}

	return strings.Join([]string{"BindSiteDto", string(data)}, " ")
}
