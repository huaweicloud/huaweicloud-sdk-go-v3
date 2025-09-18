package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AreaBandwidthPackageSpecification struct {
	LocalAreaId *LocalAreaIdDef `json:"local_area_id"`

	RemoteAreaId *RemoteAreaIdDef `json:"remote_area_id"`

	// 互通大区带宽包的规格ID。
	Id string `json:"id"`

	// 带宽包产品规格列表。
	SpecCodes []SpecificationCodeInfo `json:"spec_codes"`
}

func (o AreaBandwidthPackageSpecification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AreaBandwidthPackageSpecification struct{}"
	}

	return strings.Join([]string{"AreaBandwidthPackageSpecification", string(data)}, " ")
}
