package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkInfoCreate 切换云堡垒机实例vpc所需网络信息。
type NetworkInfoCreate struct {

	// 待创建云服务器所属虚拟私有云（简称VPC），需要指定已创建VPC的ID，UUID格式。  VPC的ID可以从控制台或者参考《虚拟私有云接口参考》的“查询VPC”章节获取。  例如：03211ecf-xxxx-4306-xxxx-6e939bfxxxxx
	VpcId string `json:"vpc_id"`

	// 子网ID，字母数字下划线连接符组成。
	SubnetId string `json:"subnet_id"`

	PublicIp *PublicIp `json:"public_ip"`

	// 安全组信息。
	SecurityGroups []SecurityGroup `json:"security_groups"`

	PrivateIp *PrivateIp `json:"private_ip,omitempty"`
}

func (o NetworkInfoCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkInfoCreate struct{}"
	}

	return strings.Join([]string{"NetworkInfoCreate", string(data)}, " ")
}
