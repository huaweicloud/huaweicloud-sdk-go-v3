package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerGroupReq 创建服务器组请求。
type CreateServerGroupReq struct {

	// 服务器组名称，名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成，不能有空格。 2. 长度范围1~64个字符。
	Name string `json:"name"`

	// 服务器组关联的镜像ID，用于创建对应组下的云服务器。
	ImageId string `json:"image_id"`

	// 服务器组的镜像产品ID，当镜像为云市场镜像时，该字段必填。
	ImageProductId *string `json:"image_product_id,omitempty"`

	ImageType *ImageTypeEnum `json:"image_type"`

	OsType *OsTypeEnum `json:"os_type"`

	// 服务器组描述。
	Description *string `json:"description,omitempty"`

	RoutePolicy *RoutePolicy `json:"route_policy,omitempty"`

	// 产品ID。 > - 获取方式详见产品套餐管理ListProduct：\"GET  /v1/{project_id}/product\"。
	ProductId string `json:"product_id"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 网卡对应的子网ID。
	SubnetId string `json:"subnet_id"`

	SystemDiskType *VolumeType `json:"system_disk_type"`

	// 磁盘容量，单位GB。
	SystemDiskSize int32 `json:"system_disk_size"`

	// 默认组织名称。
	OuName *string `json:"ou_name,omitempty"`

	// 云服务器系统盘对应的存储池的ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 可用分区。 > - 将服务创建到指定的可用分区，如果不指定则使用系统随机的可用分区。 > - 获取方式详见可用区管理ListAvailabilityZone：\"GET  /v1/{project_id}/availability-zone\"。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	IpVirtual *IpVirtual `json:"ip_virtual,omitempty"`

	// 是否为vdi单会话模式。
	IsVdi *bool `json:"is_vdi,omitempty"`

	AppType *AppTypeEnum `json:"app_type,omitempty"`

	ExtraSessionType *ExtraSessionTypeEnum `json:"extra_session_type,omitempty"`

	// 付费会话数，单位/个。
	ExtraSessionSize *int32 `json:"extra_session_size,omitempty"`

	// 标签信息，最多包含20个key,不允许重复
	Tags *[]TmsTag `json:"tags,omitempty"`

	// 企业项目ID,仅企业项目需配置(字段为空或者0表示使用默认default企业项目)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 主服务器组id,绑定主服务器组，则创建的是备服务器。
	PrimaryServerGroupId *string `json:"primary_server_group_id,omitempty"`

	// 是否启用服务器组。
	ServerGroupStatus *bool `json:"server_group_status,omitempty"`
}

func (o CreateServerGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupReq struct{}"
	}

	return strings.Join([]string{"CreateServerGroupReq", string(data)}, " ")
}
