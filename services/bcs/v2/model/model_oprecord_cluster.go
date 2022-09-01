package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群信息
type OprecordCluster struct {

	// 集群类型
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type"`

	// 集群名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id"`

	// 节点信息
	NodeInfos *[]NodeInfo `json:"node_infos,omitempty" xml:"node_infos"`
}

func (o OprecordCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OprecordCluster struct{}"
	}

	return strings.Join([]string{"OprecordCluster", string(data)}, " ")
}
