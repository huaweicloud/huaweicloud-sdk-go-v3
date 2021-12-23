package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type EdgeImageRegionInfo struct {
	// 区域ID

	RegionId *string `json:"region_id,omitempty"`
	// 镜像ID

	ImageId *string `json:"image_id,omitempty"`
}

func (o EdgeImageRegionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeImageRegionInfo struct{}"
	}

	return strings.Join([]string{"EdgeImageRegionInfo", string(data)}, " ")
}
