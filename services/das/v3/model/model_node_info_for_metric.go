package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeInfoForMetric 节点信息
type NodeInfoForMetric struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID列表
	NodeIds []string `json:"node_ids"`
}

func (o NodeInfoForMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInfoForMetric struct{}"
	}

	return strings.Join([]string{"NodeInfoForMetric", string(data)}, " ")
}
