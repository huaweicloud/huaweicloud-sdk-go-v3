package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDetailNetwork 网络信息。
type InstanceDetailNetwork struct {

	// 云堡垒机实例浮动ip。(实例为主备模式时返回对应的值)
	Vip *string `json:"vip,omitempty"`

	// 云堡垒机实例WEB界面访问的端口号。
	WebPort string `json:"web_port"`

	// 云堡垒机实例弹性公网IP。
	PublicIp string `json:"public_ip"`

	// 云堡垒机实例绑定公网的弹性IP的ID，UUID格式表示。
	PublicId string `json:"public_id"`

	// 云堡垒机实例私有ip。
	PrivateIp string `json:"private_ip"`

	// 云堡垒机实例所在虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 云堡垒机实例所在子网ID。
	SubnetId string `json:"subnet_id"`

	// 云堡垒机实例所属的安全组ID。
	SecurityGroupId string `json:"security_group_id"`
}

func (o InstanceDetailNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetailNetwork struct{}"
	}

	return strings.Join([]string{"InstanceDetailNetwork", string(data)}, " ")
}
