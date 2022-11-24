package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置服务器自动分配弹性IP时，创建弹性IP的配置参数，详情见 eip 结构体
type CreateNet2CloudPhoneServerRequestBodyPublicIpEip struct {

	// 弹性公网IP的类型，取值范围：5_telcom（电信），5_union（联通），5_bgp（全动态BGP），5_sbgp（静态BGP）
	Type string `json:"type"`
}

func (o CreateNet2CloudPhoneServerRequestBodyPublicIpEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNet2CloudPhoneServerRequestBodyPublicIpEip struct{}"
	}

	return strings.Join([]string{"CreateNet2CloudPhoneServerRequestBodyPublicIpEip", string(data)}, " ")
}
