package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nic 租户自定义的网卡的结构体，为待创建的云服务器的网卡信息。
type Nic struct {

	// 租户自定义的子网 ID，为待创建的云服务器所属的子网。  需要指定tenant_vpc_id对应VPC下已创建的子网（subnet）的网络ID，UUID格式。
	SubnetId string `json:"subnet_id"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
