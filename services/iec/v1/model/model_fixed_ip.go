package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 私有IP对象
type FixedIp struct {
	// 所属子网ID

	SubnetId *string `json:"subnet_id,omitempty"`
	// 端口IP地址

	IpAddress *string `json:"ip_address,omitempty"`
}

func (o FixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FixedIp struct{}"
	}

	return strings.Join([]string{"FixedIp", string(data)}, " ")
}
