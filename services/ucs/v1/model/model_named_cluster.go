package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamedCluster 集群信息
type NamedCluster struct {

	// 集群名称
	Name *string `json:"name,omitempty"`

	Cluster *ClusterCert `json:"cluster,omitempty"`
}

func (o NamedCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamedCluster struct{}"
	}

	return strings.Join([]string{"NamedCluster", string(data)}, " ")
}
