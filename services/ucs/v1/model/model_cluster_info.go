package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterInfo struct {

	// 集群ID
	ClusterID *string `json:"clusterID,omitempty"`

	// 集群名称
	ClusterName *string `json:"clusterName,omitempty"`

	// 集群路径
	ClusterPath *string `json:"clusterPath,omitempty"`
}

func (o ClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfo", string(data)}, " ")
}
