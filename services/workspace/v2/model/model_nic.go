package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nic struct {

	// 网卡对应的子网ID。
	SubnetId string `json:"subnet_id"`

	// IP地址，不填或空字符串将自动分配。
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
