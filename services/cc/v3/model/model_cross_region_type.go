package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CrossRegionType 跨地域类型。
type CrossRegionType struct {
	CrossRegionType *CrossRegionTypeEnum `json:"cross_region_type"`
}

func (o CrossRegionType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CrossRegionType struct{}"
	}

	return strings.Join([]string{"CrossRegionType", string(data)}, " ")
}
