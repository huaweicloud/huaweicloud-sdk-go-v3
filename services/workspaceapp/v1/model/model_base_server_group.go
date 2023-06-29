package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseServerGroup 服务器组
type BaseServerGroup struct {

	// 服务器组的唯一标识
	Id *string `json:"id,omitempty"`

	// 服务器组名称
	Name *string `json:"name,omitempty"`

	// 服务器组描述
	Description *string `json:"description,omitempty"`

	// 服务器组关联的镜像ID，用于创建对应组下的云服务器
	ImageId *string `json:"image_id,omitempty"`

	OsType *OsTypeEnum `json:"os_type,omitempty"`

	// 产品id
	ProductId *string `json:"product_id,omitempty"`

	// 网卡对应的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	SystemDiskType *VolumeType `json:"system_disk_type,omitempty"`

	// 磁盘容量，单位GB
	SystemDiskSize *int32 `json:"system_disk_size,omitempty"`

	// 是否为vdi单会话模式
	IsVdi *bool `json:"is_vdi,omitempty"`

	// 服务器组创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 服务器组更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o BaseServerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseServerGroup struct{}"
	}

	return strings.Join([]string{"BaseServerGroup", string(data)}, " ")
}
