package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecClusterNode 微服务引擎的CCE节点列表
type SpecClusterNode struct {

	// CCE节点信息。
	ClusterNodes *[]ClusterNode `json:"clusterNodes,omitempty"`
}

func (o SpecClusterNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecClusterNode struct{}"
	}

	return strings.Join([]string{"SpecClusterNode", string(data)}, " ")
}
