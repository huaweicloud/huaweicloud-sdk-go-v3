package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RocketMqDetail struct {

	// RocketMQ实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 消费组
	Group string `json:"group"`

	// Topic
	Topic string `json:"topic"`

	// 用户名
	AccessKey string `json:"access_key"`

	// 密钥
	SecretKey string `json:"secret_key"`

	// 虚拟私有云
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网
	SubnetId *string `json:"subnet_id,omitempty"`

	// 连接地址
	NamesrvAddress *string `json:"namesrv_address,omitempty"`

	// SSL
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// ACL访问控制
	EnableAcl *bool `json:"enable_acl,omitempty"`
}

func (o RocketMqDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqDetail struct{}"
	}

	return strings.Join([]string{"RocketMqDetail", string(data)}, " ")
}
