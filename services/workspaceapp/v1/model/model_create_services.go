package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServices 云服务下单实体。
type CreateServices struct {

	// 可用分区。 > - 将服务创建到指定的可用分区，如果不指定则使用系统随机的可用分区。 > - 获取方式详见可用区管理ListAvailabilityZone：\"GET  /v1/{project_id}/availability-zone\"。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 数据盘。
	DataVolumes *[]Volume `json:"data_volumes,omitempty"`

	// 网卡信息，该字段当前未使用。
	Nics *[]Nic `json:"nics,omitempty"`

	// OU名称，在对接AD时使用，需提前在AD中创建OU。
	OuName *string `json:"ou_name,omitempty"`

	// 产品ID。 > - 获取方式详见产品套餐管理ListProduct：\"GET  /v1/{project_id}/product\"。
	ProductId string `json:"product_id"`

	// 规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 操作系统类型，当前仅支持Windows - Linux - Windows
	OsType *string `json:"os_type,omitempty"`

	RootVolume *Volume `json:"root_volume"`

	// 服务器组ID, 云应用创建服务组时生成。
	ServerGroupId string `json:"server_group_id"`

	// 云服务类型，云桌面固定为DEDICATED。
	ServiceType *string `json:"service_type,omitempty"`

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	// 自动开户的时候,用于LiteAs第一次开户传进来。
	VpcId string `json:"vpc_id"`

	// 服务器使用的安全组，如果不指定则默认使用服务器代理中指定的安全组。 **⚠ 警告: 预留属性，目前未使用**
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	// 是否自动升级hda版本。
	UpdateAccessAgent *bool `json:"update_access_agent,omitempty"`
}

func (o CreateServices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServices struct{}"
	}

	return strings.Join([]string{"CreateServices", string(data)}, " ")
}
