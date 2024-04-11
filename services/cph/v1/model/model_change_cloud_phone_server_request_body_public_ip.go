package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerRequestBodyPublicIp 租户自定义的弹性公网IP结构体。  配置云服务器的弹性IP信息的方式，  自动分配，需要指定新创建弹性IP的信息； 使用已有，需要指定已经购买的EIP ID。
type ChangeCloudPhoneServerRequestBodyPublicIp struct {

	// 指定已有的EIP进行服务器创建，当前只支持传入一个已有的EIP ID。
	Ids *[]string `json:"ids,omitempty"`

	Eip *CreateNet2CloudPhoneServerRequestBodyPublicIpEip `json:"eip,omitempty"`

	// Eip数量。默认为1，不需要Eip可设置为0，取值范围为0到手机IP数。
	Count *int32 `json:"count,omitempty"`
}

func (o ChangeCloudPhoneServerRequestBodyPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequestBodyPublicIp struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequestBodyPublicIp", string(data)}, " ")
}
