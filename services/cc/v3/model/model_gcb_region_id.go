package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbRegionId struct {

	// 功能说明：实例所在region，不填默认\"global\"。
	RegionId *string `json:"region_id,omitempty"`
}

func (o GcbRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbRegionId struct{}"
	}

	return strings.Join([]string{"GcbRegionId", string(data)}, " ")
}
