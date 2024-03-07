package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIp 弹性IP信息。
type PublicIp struct {

	// 弹性IP唯一标识
	Id *string `json:"id,omitempty"`

	// 弹性IP地址
	PublicIpAddress *string `json:"public_ip_address,omitempty"`
}

func (o PublicIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIp struct{}"
	}

	return strings.Join([]string{"PublicIp", string(data)}, " ")
}
