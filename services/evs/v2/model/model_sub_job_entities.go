package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子Job的响应信息。
type SubJobEntities struct {

	// 云硬盘的类型。
	VolumeType *string `json:"volume_type,omitempty" xml:"volume_type"`

	// 云硬盘的容量，单位为GB。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 云硬盘的ID。
	VolumeId *string `json:"volume_id,omitempty" xml:"volume_id"`

	// 云硬盘的名称。
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o SubJobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJobEntities struct{}"
	}

	return strings.Join([]string{"SubJobEntities", string(data)}, " ")
}
