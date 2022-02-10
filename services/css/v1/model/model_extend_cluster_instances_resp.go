package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 扩容实例。
type ExtendClusterInstancesResp struct {
	// 实例ID。

	Id *string `json:"id,omitempty"`
	// 实例名字。

	Name *string `json:"name,omitempty"`
	// 实例类型。

	Type *string `json:"type,omitempty"`
	// 实例组名。

	ShardId *string `json:"shard_id,omitempty"`
}

func (o ExtendClusterInstancesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendClusterInstancesResp struct{}"
	}

	return strings.Join([]string{"ExtendClusterInstancesResp", string(data)}, " ")
}
