package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群节点配置
type ClusterNodeConfig struct {

	// master虚拟ip
	MasterNodeVip *string `json:"master_node_vip,omitempty"`

	// master节点数
	MasterNodes *[]NodeConfig `json:"master_nodes,omitempty"`

	// work节点数
	WorkNodes *[]NodeConfig `json:"work_nodes,omitempty"`
}

func (o ClusterNodeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeConfig struct{}"
	}

	return strings.Join([]string{"ClusterNodeConfig", string(data)}, " ")
}
