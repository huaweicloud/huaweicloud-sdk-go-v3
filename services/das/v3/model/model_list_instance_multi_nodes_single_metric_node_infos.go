package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstanceMultiNodesSingleMetricNodeInfos struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID列表
	NodeIds []string `json:"node_ids"`
}

func (o ListInstanceMultiNodesSingleMetricNodeInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMultiNodesSingleMetricNodeInfos struct{}"
	}

	return strings.Join([]string{"ListInstanceMultiNodesSingleMetricNodeInfos", string(data)}, " ")
}
