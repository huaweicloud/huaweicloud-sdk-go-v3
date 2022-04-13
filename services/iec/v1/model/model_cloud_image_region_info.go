package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CloudImageRegionInfo struct {
	// 区域ID

	RegionId *string `json:"region_id,omitempty"`
	// 镜像ID

	ImageId *string `json:"image_id,omitempty"`
}

func (o CloudImageRegionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudImageRegionInfo struct{}"
	}

	return strings.Join([]string{"CloudImageRegionInfo", string(data)}, " ")
}
