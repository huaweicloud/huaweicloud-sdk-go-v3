package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nic 网络网卡信息
type Nic struct {

	// 网卡信息，如eth0,eth1
	Eth *string `json:"eth,omitempty"`

	// 网卡ip
	Ip *string `json:"ip,omitempty"`

	// 子网掩码的位数
	MaskLen *int32 `json:"mask_len,omitempty"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
