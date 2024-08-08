package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerGroupResponse Response Object
type CreateServerGroupResponse struct {

	// 服务器组的唯一标识。
	Id *string `json:"id,omitempty"`

	// 服务器组名称。
	Name *string `json:"name,omitempty"`

	// 服务器组描述。
	Description *string `json:"description,omitempty"`

	// 服务器组关联的镜像ID，用于创建对应组下的云服务器。
	ImageId *string `json:"image_id,omitempty"`

	OsType *OsTypeEnum `json:"os_type,omitempty"`

	// 产品id。
	ProductId *string `json:"product_id,omitempty"`

	// 网卡对应的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	SystemDiskType *VolumeType `json:"system_disk_type,omitempty"`

	// 磁盘容量，单位GB。
	SystemDiskSize *int32 `json:"system_disk_size,omitempty"`

	// 是否为vdi单会话模式。
	IsVdi *bool `json:"is_vdi,omitempty"`

	ExtraSessionType *ExtraSessionTypeEnum `json:"extra_session_type,omitempty"`

	// 付费会话个数。
	ExtraSessionSize *int32 `json:"extra_session_size,omitempty"`

	AppType *AppTypeEnum `json:"app_type,omitempty"`

	// 服务器组创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 服务器组更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	StorageMountPolicy *StorageFolderMountType `json:"storage_mount_policy,omitempty"`

	// 企业项目ID(0表示默认企业项目Id)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主服务器组id列表。
	PrimaryServerGroupIds *[]string `json:"primary_server_group_ids,omitempty"`

	// 备服务器组id列表。
	SecondaryServerGroupIds *[]string `json:"secondary_server_group_ids,omitempty"`

	// 服务器是否处于启用状态，true表示处于启用状态 false表示处于禁用状态。
	ServerGroupStatus *bool `json:"server_group_status,omitempty"`

	// 站点类型 - CENTER/IES
	SiteType *string `json:"site_type,omitempty"`

	// 站点id
	SiteId *string `json:"site_id,omitempty"`

	// 服务器配置总数量。
	AppServerFlavorCount *int32 `json:"app_server_flavor_count,omitempty"`

	// 服务器总数量。
	AppServerCount *int32 `json:"app_server_count,omitempty"`

	// 关联应用组的总数量。
	AppGroupCount *int32 `json:"app_group_count,omitempty"`

	// 镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty"`

	ScalingPolicy *ScalingPolicy `json:"scaling_policy,omitempty"`

	// 标签信息
	Tags *[]TmsTag `json:"tags,omitempty"`

	// 默认组织名称。
	OuName         *string `json:"ou_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateServerGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateServerGroupResponse", string(data)}, " ")
}
