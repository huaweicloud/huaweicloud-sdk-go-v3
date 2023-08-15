package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteClusterSwitchoverRequest Request Object
type ExecuteClusterSwitchoverRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 分片ID
	GroupId string `json:"group_id"`

	// 升级为主节点的节点ID
	NodeId string `json:"node_id"`
}

func (o ExecuteClusterSwitchoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterSwitchoverRequest struct{}"
	}

	return strings.Join([]string{"ExecuteClusterSwitchoverRequest", string(data)}, " ")
}
