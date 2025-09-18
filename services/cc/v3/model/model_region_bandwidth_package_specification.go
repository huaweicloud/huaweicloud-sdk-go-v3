package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionBandwidthPackageSpecification 城域带宽包实例。
type RegionBandwidthPackageSpecification struct {

	// RegionID。
	LocalRegionId string `json:"local_region_id"`

	// RegionID。
	RemoteRegionId string `json:"remote_region_id"`

	// 互通区域带宽包的规格ID。
	Id *string `json:"id,omitempty"`

	// 互通区域带宽包的规格名称。
	Name *string `json:"name,omitempty"`

	// 互通区域带宽包的规格英文名称。
	EnName *string `json:"en_name,omitempty"`

	// 互通区域带宽包的规格西语名称。
	EsName *string `json:"es_name,omitempty"`

	// 互通区域带宽包的规格葡语名称。
	PtName *string `json:"pt_name,omitempty"`

	// 带宽包产品规格列表。
	SpecCodes *[]SpecificationCodeInfo `json:"spec_codes,omitempty"`
}

func (o RegionBandwidthPackageSpecification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionBandwidthPackageSpecification struct{}"
	}

	return strings.Join([]string{"RegionBandwidthPackageSpecification", string(data)}, " ")
}
