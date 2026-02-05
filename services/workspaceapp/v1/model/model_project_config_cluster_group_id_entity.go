package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectConfigClusterGroupIdEntity 项目配置ID信息，包含项目配置ID和集群ID。
type ProjectConfigClusterGroupIdEntity struct {

	// 集群ID
	ClusterGroupId *string `json:"cluster_group_id,omitempty"`

	// 项目配置ID
	ProjectConfigId *string `json:"project_config_id,omitempty"`
}

func (o ProjectConfigClusterGroupIdEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectConfigClusterGroupIdEntity struct{}"
	}

	return strings.Join([]string{"ProjectConfigClusterGroupIdEntity", string(data)}, " ")
}
