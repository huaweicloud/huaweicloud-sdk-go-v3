package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EncryptDataNodes struct {

	// 边缘节点ID
	Id string `json:"id"`

	// 边缘节点状态
	State string `json:"state"`

	// 边缘节点名称
	Name string `json:"name"`

	// 边缘节点主机名
	HostName string `json:"host_name"`

	// 边缘节点主机IP地址列表
	HostIps []string `json:"host_ips"`
}

func (o EncryptDataNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataNodes struct{}"
	}

	return strings.Join([]string{"EncryptDataNodes", string(data)}, " ")
}
