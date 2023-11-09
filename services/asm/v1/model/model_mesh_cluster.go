package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MeshCluster struct {

	// 集群ID，资源唯一标识，通过该ID查询需要添加的集群
	ClusterID string `json:"clusterID"`

	// 集群所属的projectID
	ProjectID string `json:"projectID"`

	Injection *InjectionConfig `json:"injection,omitempty"`

	Installation *InstallationConfig `json:"installation"`
}

func (o MeshCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshCluster struct{}"
	}

	return strings.Join([]string{"MeshCluster", string(data)}, " ")
}
