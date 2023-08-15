package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseEndpointConfig 数据库基本设置信息体，当源库或者目标库在特定场景下需要额外设置时填写。例如： - 实时迁移入云场景：云数据库（RDS for MySQL）、云数据库 GaussDB(for MySQL) 主备版作为目标库（target_endpoint）时，需要设置“只读”或者“读写”（is_target_readonly)。 - 实时迁移入云场景：分布式数据库中间件DDM作为目标库，源库为MySQL分库分表时，需要设置“端数据库实例个数”。 - 实时迁移入云场景：MongoDB作为源库时，需要设置“源数据库实例类型”为非集群、集群或者集群(MongoDB 4.0+)，并且实例类型为“集群”时，还需要设置“源端分片个数”（source_shard_num)。 - 实时灾备场景：单主灾备且本云为备时，云数据库（RDS for MySQL）、云数据库 GaussDB(for MySQL) 主备版、分布式数据库中间件DDM、文档数据库服务DDS等作为目标库（target_endpoint）时，需要设置“只读”或者“读写”（is_target_readonly)。 - 实时灾备场景：云数据库GaussDB(for Cassandra) 灾备时，需要设置“是否迁移TTL”，如果参数为true时，还需要设置TTL列名。
type BaseEndpointConfig struct {

	// 目标实例是否设置为为只读。 - MySQL迁移和灾备，且job_direction为up时设置有效。（灾备场景下，单主灾备且本云为备为必填且为true，不填默认设置为true）。
	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`
}

func (o BaseEndpointConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseEndpointConfig struct{}"
	}

	return strings.Join([]string{"BaseEndpointConfig", string(data)}, " ")
}
