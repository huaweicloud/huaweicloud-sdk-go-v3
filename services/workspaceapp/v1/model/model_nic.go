package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Nic 云服务对应的网卡信息
type Nic struct {

	// IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 网卡对应的子网ID
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
