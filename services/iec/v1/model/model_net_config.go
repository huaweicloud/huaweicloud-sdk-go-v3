package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetConfig
type NetConfig struct {

	// 边缘网络ID。  约束： - 创建边缘业务仅支持使用系统规划的虚拟私有云。
	VpcId string `json:"vpc_id"`

	// 边缘实例绑定的网卡数量。  约束：一台边缘实例最少绑定一张网卡，最多绑定8张网卡。
	NicNum int32 `json:"nic_num"`

	// - 功能说明：IP/Mac对列表 - 约束：     IP地址不允许为 “0.0.0.0/0”     如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。     如果allowed_address_pairs为“1.1.1.1/0”，表示关闭源目地址检查开关
	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`
}

func (o NetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetConfig struct{}"
	}

	return strings.Join([]string{"NetConfig", string(data)}, " ")
}
