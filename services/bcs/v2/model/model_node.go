package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Node struct {
	// 节点EIP信息

	IpPort *string `json:"ip_port,omitempty"`
	// 节点所在通道数组

	Channels *[]string `json:"channels,omitempty"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
