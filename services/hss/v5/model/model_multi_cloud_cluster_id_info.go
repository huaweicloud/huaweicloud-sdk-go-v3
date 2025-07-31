package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiCloudClusterIdInfo 多云集群的集群id信息
type MultiCloudClusterIdInfo struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o MultiCloudClusterIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiCloudClusterIdInfo struct{}"
	}

	return strings.Join([]string{"MultiCloudClusterIdInfo", string(data)}, " ")
}
