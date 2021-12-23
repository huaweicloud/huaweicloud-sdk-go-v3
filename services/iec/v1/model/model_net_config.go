package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NetConfig struct {
	// 边缘网络ID。  约束： - 创建边缘业务仅支持使用系统规划的虚拟私有云。

	VpcId string `json:"vpc_id"`
	// 边缘实例绑定的网卡数量。  约束：一台边缘实例最少绑定一张网卡，最多绑定8张网卡。

	NicNum int32 `json:"nic_num"`
}

func (o NetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetConfig struct{}"
	}

	return strings.Join([]string{"NetConfig", string(data)}, " ")
}
