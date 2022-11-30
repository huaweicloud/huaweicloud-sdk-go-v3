package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeploymentFormResponse struct {

	// 初始节点数。
	InitialNodeNum *int32 `json:"initial_node_num,omitempty"`

	// 解决方案模板名称。
	Solution *string `json:"solution,omitempty"`

	// 分片数。
	ShardNum *int32 `json:"shard_num,omitempty"`

	// 副本数。
	ReplicaNum     *int32 `json:"replica_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeploymentFormResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentFormResponse struct{}"
	}

	return strings.Join([]string{"ShowDeploymentFormResponse", string(data)}, " ")
}
