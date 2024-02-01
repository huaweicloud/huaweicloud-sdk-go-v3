package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivateIp 云堡垒机实例指定公网IP信息。
type PrivateIp struct {

	// 私网IP地址。
	Ip string `json:"ip"`

	// 备机私网IP地址。
	SlaveIp *string `json:"slave_ip,omitempty"`

	// 浮动IP地址。
	FloatingIp *string `json:"floating_ip,omitempty"`
}

func (o PrivateIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateIp struct{}"
	}

	return strings.Join([]string{"PrivateIp", string(data)}, " ")
}
