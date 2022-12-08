package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建堡垒机实例请求body
type CbhInstances struct {

	// 镜像ID
	ImageRef *string `json:"image_ref,omitempty"`

	// 规格ID
	FlavorRef string `json:"flavor_ref"`

	// 堡垒机实例名称
	InstanceName string `json:"instance_name"`

	// 名字
	Name *string `json:"name,omitempty"`

	Personality *Personality `json:"personality,omitempty"`

	// 注入用户数据
	UserData *string `json:"user_data,omitempty"`

	// 初始登录密码
	AdminPassword *string `json:"admin_password,omitempty"`

	// 管理员SSH秘钥登录
	KeyName *string `json:"key_name,omitempty"`

	// VPC ID
	VpcId string `json:"vpc_id"`

	// 网卡信息
	Nics []Nics `json:"nics"`

	PublicIp *PublicIp `json:"public_ip"`

	// 弹性数量
	Count *int32 `json:"count,omitempty"`

	RootVolume *RootVolume `json:"root_volume,omitempty"`

	DataVolumes *DataVolumes `json:"data_volumes,omitempty"`

	// 网卡信息
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 分区信息
	AvailabilityZone string `json:"availability_zone"`

	// 备用区
	SlaveAvailabilityZone string `json:"slave_availability_zone"`

	ExtendParam *ExtendParam `json:"extend_param,omitempty"`

	// 创建云服务云数据
	Metadata *string `json:"metadata,omitempty"`

	// 描述
	Comment *string `json:"comment,omitempty"`

	// 云服务所在区域ID
	Region string `json:"region"`

	// region标识
	RegionId *string `json:"region_id,omitempty"`

	// 资源规格
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 前端登录密码
	HxPassword string `json:"hx_password"`

	// 堡垒机机机型
	BastionType string `json:"bastion_type"`

	// 分区信息
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 订购截止日期
	EndTime *string `json:"end_time,omitempty"`
}

func (o CbhInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbhInstances struct{}"
	}

	return strings.Join([]string{"CbhInstances", string(data)}, " ")
}
