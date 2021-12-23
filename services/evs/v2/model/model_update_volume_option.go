package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateVolumeOption struct {
	// 新的云硬盘的描述，name和description不能同时为null。最大支持255个字节。

	Description *string `json:"description,omitempty"`
	// 新的云硬盘的名字，name和description不能同时为null。最大支持255个字节。

	Name *string `json:"name,omitempty"`
}

func (o UpdateVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeOption struct{}"
	}

	return strings.Join([]string{"UpdateVolumeOption", string(data)}, " ")
}
