package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群信息
type OprecordCluster struct {
	// 集群类型

	ClusterType *string `json:"cluster_type,omitempty"`
	// 集群名称

	Name *string `json:"name,omitempty"`
	// 集群ID

	ClusterId *string `json:"cluster_id,omitempty"`
	// 节点信息

	NodeInfos *[]NodeInfo `json:"node_infos,omitempty"`
}

func (o OprecordCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OprecordCluster struct{}"
	}

	return strings.Join([]string{"OprecordCluster", string(data)}, " ")
}
