package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IefNodeinfo struct {
	// 节点状态，要使用此节点的话，该状态值必须为ACTIVE

	Status string `json:"status"`
	// 节点IP，填写节点所在的EIP地址

	PublicIpAddress string `json:"public_ip_address"`
	// ief节点id值

	Id string `json:"id"`
}

func (o IefNodeinfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IefNodeinfo struct{}"
	}

	return strings.Join([]string{"IefNodeinfo", string(data)}, " ")
}
