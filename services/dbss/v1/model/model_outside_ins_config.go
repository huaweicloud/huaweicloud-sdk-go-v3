package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutsideInsConfig 云实例配置
type OutsideInsConfig struct {

	// 主节点IP
	MasterNodeIp *string `json:"master_node_ip,omitempty"`

	// 从节点IP
	SlaveNodeIp *string `json:"slave_node_ip,omitempty"`

	// 虚拟IP
	VirtualIp *string `json:"virtual_ip,omitempty"`
}

func (o OutsideInsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutsideInsConfig struct{}"
	}

	return strings.Join([]string{"OutsideInsConfig", string(data)}, " ")
}
