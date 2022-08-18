package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNet2CloudPhoneServerRequestBodyNics struct {

	// 租户自定义的子网 ID，为待创建的云服务器所属的子网。  需要指定tenant_vpc_id对应VPC下已创建的子网（subnet）的网络ID，UUID格式
	SubnetId string `json:"subnet_id"`
}

func (o CreateNet2CloudPhoneServerRequestBodyNics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyNics struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyNics", string(data)}, " ")
}
