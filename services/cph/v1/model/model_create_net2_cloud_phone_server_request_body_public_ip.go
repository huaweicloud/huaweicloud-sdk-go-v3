package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 租户自定义的弹性公网IP结构体。  配置云服务器的弹性IP信息的方式为：  自动分配，需要指定新创建弹性IP的信息 使用已有，需要指定已经购买的EIP ID
type CreateNet2CloudPhoneServerRequestBodyPublicIp struct {

	// 指定已有的EIP进行服务器创建，当前只支持传入一个已有的EIP ID
	Ids *[]string `json:"ids,omitempty"`

	Eip *CreateNet2CloudPhoneServerRequestBodyPublicIpEip `json:"eip,omitempty"`
}

func (o CreateNet2CloudPhoneServerRequestBodyPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyPublicIp struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyPublicIp", string(data)}, " ")
}
