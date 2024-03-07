package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeDetail 磁盘信息。
type VolumeDetail struct {

	// 桌面数据盘对应的磁盘类型，需要与系统所提供的磁盘类型相匹配。  - SAS：高IO。 - SSD：超高IO。
	Type string `json:"type"`

	// 磁盘容量，单位GB。
	Size int32 `json:"size"`

	// 挂载目录。
	Device *string `json:"device,omitempty"`

	// 磁盘表唯一标识ID。
	Id *string `json:"id,omitempty"`

	// 磁盘ID。
	VolumeId *string `json:"volume_id,omitempty"`

	// 磁盘计费资源ID。
	BillResourceId *string `json:"bill_resource_id,omitempty"`

	// 磁盘的创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 磁盘名
	DisplayName *string `json:"display_name,omitempty"`
}

func (o VolumeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDetail struct{}"
	}

	return strings.Join([]string{"VolumeDetail", string(data)}, " ")
}
