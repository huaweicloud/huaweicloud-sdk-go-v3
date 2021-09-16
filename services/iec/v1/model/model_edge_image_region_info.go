package model

import (
	"encoding/json"

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
	data, err := json.Marshal(o)
	if err != nil {
		return "EdgeImageRegionInfo struct{}"
	}

	return strings.Join([]string{"EdgeImageRegionInfo", string(data)}, " ")
}
