package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeList struct {

	// 端口。
	Port *string `json:"port,omitempty" xml:"port"`

	// 节点状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 节点id。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// ip
	Ip *string `json:"ip,omitempty" xml:"ip"`
}

func (o NodeList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeList struct{}"
	}

	return strings.Join([]string{"NodeList", string(data)}, " ")
}
