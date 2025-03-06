package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneSingleServerRequestBodyPublicIp 租户自定义的弹性公网IP结构体。 配置云手机裸服务器的弹性IP信息的方式， 自动分配，需要指定新创建弹性IP的信息； 使用已有，需要指定已经购买的EIP ID。
type CreateCloudPhoneSingleServerRequestBodyPublicIp struct {

	// 指定已有的EIP进行云手机裸服务器创建，当前只支持传入一个已有的EIP ID。指定EIP后public_ip结构体中count字段和type字段不生效。
	Ids *[]string `json:"ids,omitempty"`

	// 弹性公网IP的类型。 例如： 5_telcom：电信 5_union：联通 5_bgp：全动态BGP 5_sbgp：静态BGP
	Type *string `json:"type,omitempty"`

	// Eip数量。默认为0。取值范围为【0,1】
	Count *int32 `json:"count,omitempty"`

	BandWidth *CreateCloudPhoneSingleServerRequestBodyPublicIpBandWidth `json:"band_width"`
}

func (o CreateCloudPhoneSingleServerRequestBodyPublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneSingleServerRequestBodyPublicIp struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneSingleServerRequestBodyPublicIp", string(data)}, " ")
}
