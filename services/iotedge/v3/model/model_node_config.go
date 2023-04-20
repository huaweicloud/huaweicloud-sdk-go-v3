package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群节点信息
type NodeConfig struct {

	// 节点IP
	NodeIp *string `json:"node_ip,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`
}

func (o NodeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeConfig struct{}"
	}

	return strings.Join([]string{"NodeConfig", string(data)}, " ")
}
