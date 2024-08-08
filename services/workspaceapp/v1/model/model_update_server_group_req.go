package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerGroupReq 更新服务器组请求。
type UpdateServerGroupReq struct {

	// 服务器组名称，名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成，不能有空格。 2. 长度范围1~64个字符。
	Name *string `json:"name,omitempty"`

	// 服务器组描述。
	Description *string `json:"description,omitempty"`

	RoutePolicy *RoutePolicy `json:"route_policy,omitempty"`

	StorageMountPolicy *StorageFolderMountType `json:"storage_mount_policy,omitempty"`

	// 服务器组关联的镜像ID，更新镜像ID只对组下新创建的云服务器生效。
	ImageId *string `json:"image_id,omitempty"`

	// 服务器组的镜像的productId。
	ImageProductId *string `json:"image_product_id,omitempty"`

	ImageType *ImageTypeEnum `json:"image_type,omitempty"`

	SystemDiskType *VolumeType `json:"system_disk_type,omitempty"`

	// 磁盘容量，单位GB。
	SystemDiskSize *int32 `json:"system_disk_size,omitempty"`

	// 默认组织名称。
	OuName *string `json:"ou_name,omitempty"`

	AppType *AppTypeEnum `json:"app_type,omitempty"`

	// 服务器是否处于禁用状态： * `true` - 启用状态 * `false` - 禁用状态
	ServerGroupStatus *bool `json:"server_group_status,omitempty"`
}

func (o UpdateServerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerGroupReq struct{}"
	}

	return strings.Join([]string{"UpdateServerGroupReq", string(data)}, " ")
}
