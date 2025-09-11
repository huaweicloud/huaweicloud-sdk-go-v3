package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nic struct {

	// IP地址，不填或空字符串为自动分配
	IpAddress *string `json:"ip_address,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}
