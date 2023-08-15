package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudVpcInfo 数据库实例所在VPC，子网，安全组等信息，当数据库实例类型为ecs（华为云ECS自建数据库），cloud（华为云数据库）时为必填项。
type CloudVpcInfo struct {

	// 数据库实例所在的虚拟私有云ID，获取方法如下： 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询VPC列表。
	VpcId string `json:"vpc_id"`

	// 数据库实例所在子网ID，获取方法如下： 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询子网列表。
	SubnetId string `json:"subnet_id"`

	// 数据库实例所在的安全组ID，获取方法如下： 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考查询安全组列表。
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o CloudVpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudVpcInfo struct{}"
	}

	return strings.Join([]string{"CloudVpcInfo", string(data)}, " ")
}
