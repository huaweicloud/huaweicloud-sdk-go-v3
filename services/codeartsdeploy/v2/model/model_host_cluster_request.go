package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostClusterRequest struct {

	// 主机集群名称
	Name string `json:"name"`

	// 主机集群描述
	Description *string `json:"description,omitempty"`

	// slave集群id，默认为null时使用八爪鱼slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`
}

func (o HostClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostClusterRequest struct{}"
	}

	return strings.Join([]string{"HostClusterRequest", string(data)}, " ")
}
