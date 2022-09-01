package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 逻辑卷组信息
type VolumeGroups struct {

	// Pv信息
	Components *string `json:"components,omitempty" xml:"components"`

	// 剩余空间
	FreeSize *int64 `json:"free_size,omitempty" xml:"free_size"`

	// lv信息
	LogicalVolumes *[]LogicalVolumes `json:"logical_volumes,omitempty" xml:"logical_volumes"`

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 大小
	Size *int64 `json:"size,omitempty" xml:"size"`
}

func (o VolumeGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeGroups struct{}"
	}

	return strings.Join([]string{"VolumeGroups", string(data)}, " ")
}
