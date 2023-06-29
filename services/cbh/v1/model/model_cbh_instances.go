package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbhInstances 创建堡垒机实例请求参数。
type CbhInstances struct {

	// 待创建云堡垒机规格ID，例如： - cbh.basic.50 - cbh.enhance.50 已上线的规格请参见《云堡垒机产品介绍》的[服务版本差异](https://support.huaweicloud.com/productdesc-cbh/cbh_01_0010.html)章节。
	FlavorRef string `json:"flavor_ref"`

	// 云堡垒机实例名称，取值范围： - 只能由中文字符、英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。 例如：CBH-6b8e
	InstanceName string `json:"instance_name"`

	// 待创建云服务器所属虚拟私有云（简称VPC），需要指定已创建VPC的ID，UUID格式。 VPC的ID可以从控制台或者参考《虚拟私有云接口参考》的“查询VPC”章节获取。 例如：03211ecf-697e-4306-a7a0-6e939bf948de
	VpcId string `json:"vpc_id"`

	Nics []Nics `json:"nics"`

	PublicIp *PublicIp `json:"public_ip"`

	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 创建云堡垒机所在的可用区，需要指定可用区名称。 可参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)获取
	AvailabilityZone string `json:"availability_zone"`

	// 创建云堡垒机所在的备机可用区，需要指定备机可用区名称。(当前字段未启用,填写默认值null) 可参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)获取
	SlaveAvailabilityZone *string `json:"slave_availability_zone,omitempty"`

	// 云堡垒机实例描述信息。
	Comment *string `json:"comment,omitempty"`

	// 云服务所在局点ID。
	Region string `json:"region"`

	// 堡垒机实例前端登录密码。密码规则：8-32位,不能包含amdin或nidma及其大写形式,必须包含大小写数字特殊字符四种类型中的三种。
	HxPassword string `json:"hx_password"`

	// 堡垒机实例类型，填写“OEM”即可。
	BastionType string `json:"bastion_type"`

	// 是否支持IPV6，不填默认为false。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 企业项目ID，不填默认为0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CbhInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbhInstances struct{}"
	}

	return strings.Join([]string{"CbhInstances", string(data)}, " ")
}
