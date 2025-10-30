package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostStatusInfo struct {

	// **参数解释**: 服务器ID **取值范围**: 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-128位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 弹性公网IP地址
	PublicIp *[]string `json:"public_ip,omitempty"`

	// **参数解释**: 私有IP地址
	PrivateIp *[]string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: 字符长度1-128位
	Status *string `json:"status,omitempty"`

	// **参数解释**: Vpc标识Id **取值范围**: 字符长度1-128位
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o HostStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostStatusInfo struct{}"
	}

	return strings.Join([]string{"HostStatusInfo", string(data)}, " ")
}
