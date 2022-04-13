package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeType struct {
	// 云硬盘类型的ID。

	Id string `json:"id"`
	// 云硬盘类型名称。

	Name string `json:"name"`

	ExtraSpecs *VolumeTypeExtraSpecs `json:"extra_specs,omitempty"`
	// 云硬盘类型的描述信息。

	Description *string `json:"description,omitempty"`
	// 预留属性。

	QosSpecsId *string `json:"qos_specs_id,omitempty"`
	// 预留属性。

	IsPublic *bool `json:"is_public,omitempty"`
}

func (o VolumeType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeType struct{}"
	}

	return strings.Join([]string{"VolumeType", string(data)}, " ")
}
